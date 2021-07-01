package gm

// Tile 地图瓦片
type Tile struct {
	X, Y, Z int
}

/**
 * 返回当前坐标上移 distance 行对应的坐标，即：Row - distance
 * @param	distance	往上的行数
 * @return			对应坐标
 */
func (m Tile) Up(distance int) *Tile {
	return &Tile{m.Y - distance, m.X, m.Z}
}

/**
 * 返回当前坐标右移 distance 列对应的坐标，即：Column + distance
 * @param	distance	往右的列数
 * @return			对应坐标
 */
func (m Tile) Right(distance int) *Tile { // = 1
	return &Tile{m.Y, m.X + distance, m.Z}
}

/**
 * 返回当前坐标下移 distance 行对应的坐标，即：Row + distance
 * @param	distance	下移的行数
 * @return			对应坐标
 */
func (m Tile) Down(distance int) *Tile { //
	return &Tile{m.Y + distance, m.X, m.Z}
}

/**
 * 返回当前坐标左移 distance 列对应的坐标，即：Column - distance
 * @param	distance	往左的列数
 * @return			对应坐标
 */
func (m Tile) Left(distance int) *Tile {
	return &Tile{m.Y, m.X - distance, m.Z}
}
