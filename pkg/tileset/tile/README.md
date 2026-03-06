# tile

## GeometricError
GeometricError 是 3D Tiles 数据格式中的一个重要参数，用于描述某个 Tile 的简化几何体与原始几何体之间的偏差。它是一个非负数，单位为米

屏幕空间误差(SSE)的计算公式

SSE = (geometricError×height) / (distance×sseDenominator)

- height 是屏幕的像素高度
- distance 是相机到 Tile 的距离。
- sseDenominator 是一个根据视锥体的张角和宽高比计算的参数

SSE 超过某个阈值（如 1 像素），则加载更详细的子 Tile

经验上看：

geometricError≈distance×0.00957
