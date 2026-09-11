package udp

import (
	"errors"
	"fmt"
	"net"

	"github.com/daanv2/race-game-dashboard/pkg/extensions/xsync"
)

// MaxDatagramSize is large enough to hold the biggest possible UDP payload.
// A datagram larger than the buffer passed to ReadFromUDP is silently
// truncated by the OS, so this must be >= the largest packet you expect.
const MaxDatagramSize = 65535 // 65KB
const DefaultByteSize = 4096  // 4kb

type Server struct {
	conn     *net.UDPConn
	bytePool *xsync.Pool[[]byte]

	Handler Handler
}

func NewServer(port uint16, handler Handler) (*Server, error) {
	addr := fmt.Sprintf(":%d", port)
	udpaddr, err := net.ResolveUDPAddr("udp", addr)
	if err != nil {
		return nil, fmt.Errorf("error resolving udp addr: %w", err)
	}

	conn, err := net.ListenUDP("udp", udpaddr)
	if err != nil {
		return nil, fmt.Errorf("listing on udp addr '%s': %w", addr, err)
	}

	return &Server{
		conn: conn,
		bytePool: xsync.NewPool(func() []byte {
			return make([]byte, DefaultByteSize)
		}),
		Handler: handler,
	}, nil
}

func (s *Server) Listen() (err error) {
	buf := make([]byte, MaxDatagramSize)
	var n int

	for {
		n, err = s.conn.Read(buf)
		if err != nil {
			s.bytePool.Put(buf)

			if errors.Is(err, net.ErrClosed) {
				return nil
			}

			fmt.Print(fmt.Errorf("reading udp packet: %w", err))
		} else {
			// Free up buf, use another buffer and pass that
			callBuff := s.bytePool.Get()
			if n > len(callBuff) { // Too small, make a new one, and throw out the old one that fits
				callBuff = make([]byte, n)
			}
			callBuff = callBuff[:n]
			copy(callBuff[:n], buf[:n])

			// Async handle it
			go s.sendToHandlerWithCleanup(callBuff, n)
		}
	}
}

func (s *Server) Close() error {
	return s.conn.Close()
}

// sendToHandlerWithCleanup send the given data to the handler, and when done, or crashed, hands back the data to the pool
func (s *Server) sendToHandlerWithCleanup(data []byte, n int) {
	defer func() {
		s.bytePool.Put(data)
	}()

	if s.Handler != nil {
		s.Handler.HandleMsg(data[:n])
	}
}
