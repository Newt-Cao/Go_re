module vivianchan.cn/Basics/src/Demo_twelve/customerInformation/main

go 1.17

require customerView v0.0.0-00010101000000-000000000000

require (
	customerModule v0.0.0-00010101000000-000000000000 // indirect
	customerService v0.0.0 // indirect
)

replace (
	customerModule => ../customerModule
	customerService => ../customerService
	customerView => ../customerView
)
