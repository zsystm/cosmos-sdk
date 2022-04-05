package ormtable

type GasMeter interface {
	ConsumeReadGas(keyBytes, valueBytes int)
	ConsumeMarshalGas(nBytes int)
}
