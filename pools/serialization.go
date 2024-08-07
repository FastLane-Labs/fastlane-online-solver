package pools

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
)

func SerializePools(pools []Pool) ([]byte, error) {
	buf := new(bytes.Buffer)
	for _, pool := range pools {
		poolType := pool.PoolType()
		if err := binary.Write(buf, binary.LittleEndian, poolType); err != nil {
			return nil, err
		}
		serializedPool, err := pool.Serialize()
		if err != nil {
			return nil, err
		}
		length := uint16(len(serializedPool))
		if err := binary.Write(buf, binary.LittleEndian, length); err != nil {
			return nil, err
		}
		if _, err := buf.Write(serializedPool); err != nil {
			return nil, err
		}
	}
	return buf.Bytes(), nil
}

func DeserializePools(data []byte) ([]Pool, error) {
	var pools []Pool
	buf := bytes.NewReader(data)
	for buf.Len() > 0 {
		var poolType PoolType
		if err := binary.Read(buf, binary.LittleEndian, &poolType); err != nil {
			return nil, err
		}
		var length uint16
		if err := binary.Read(buf, binary.LittleEndian, &length); err != nil {
			return nil, err
		}
		serializedPool := make([]byte, length)
		if _, err := buf.Read(serializedPool); err != nil {
			return nil, err
		}
		var pool Pool
		switch poolType {
		case UniswapV2PoolType:
			pool = &UniswapV2Pool{}
		case UniswapV3PoolType:
			pool = &UniswapV3Pool{}
		default:
			return nil, fmt.Errorf("unknown pool type: %d", poolType)
		}
		if err := pool.Deserialize(serializedPool); err != nil {
			return nil, err
		}
		pools = append(pools, pool)
	}
	return pools, nil
}

func SerializePoolsToFile(filename string, pools []Pool) error {
	data, err := SerializePools(pools)
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0644)
}

func DeserializePoolsFromFile(filename string) ([]Pool, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return DeserializePools(data)
}

func SerializeSwapPaths(paths []*SwapPath) ([]byte, error) {
	buf := new(bytes.Buffer)
	for _, path := range paths {
		serializedPath, err := path.Serialize()
		if err != nil {
			return nil, err
		}
		length := uint16(len(serializedPath))
		if err := binary.Write(buf, binary.LittleEndian, length); err != nil {
			return nil, err
		}
		if _, err := buf.Write(serializedPath); err != nil {
			return nil, err
		}
	}
	return buf.Bytes(), nil
}

func DeserializeSwapPaths(data []byte) ([]*SwapPath, error) {
	var paths []*SwapPath
	buf := bytes.NewReader(data)
	for buf.Len() > 0 {
		var length uint16
		if err := binary.Read(buf, binary.LittleEndian, &length); err != nil {
			return nil, err
		}
		serializedPath := make([]byte, length)
		if _, err := buf.Read(serializedPath); err != nil {
			return nil, err
		}
		path := &SwapPath{}
		if err := path.Deserialize(serializedPath); err != nil {
			return nil, err
		}
		paths = append(paths, path)
	}
	return paths, nil
}

func SerializeSwapPathsToFile(filename string, paths []*SwapPath) error {
	data, err := SerializeSwapPaths(paths)
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0644)
}

func DeserializeSwapPathsFromFile(filename string) ([]*SwapPath, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return DeserializeSwapPaths(data)
}
