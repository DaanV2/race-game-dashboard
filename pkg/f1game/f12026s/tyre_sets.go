package f12026s

type TyreSetData struct {
	m_actualTyreCompound uint8
	m_visualTyreCompound uint8
	m_wear               uint8
	m_available          uint8
	m_recommendedSession uint8
	m_lifeSpan           uint8
	m_usableLife         uint8
	m_lapDeltaTime       int16
	m_fitted             uint8
}

type PacketTyreSetsData struct {
	m_header      PacketHeader
	m_carIdx      uint8       // Index of the car this data relates to
	m_tyreSetData TyreSetData //  13 (dry) + 7 (wet)
	m_fittedIdx   uint8       // Index into array of fitted tyre
}
