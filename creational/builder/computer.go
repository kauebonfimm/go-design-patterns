package builder

type Component struct {
	Name  string
	Price float64
}

type CPU struct {
	Component
}
type Memory struct {
	Component
}
type Storage struct {
	Component
}
type Motherboard struct {
	Component
}

type Computer struct {
	cpu         CPU
	memory      Memory
	storage     Storage
	motherboard Motherboard
}

func (c *Computer) Price() float64 {
	return c.cpu.Price + c.memory.Price + c.storage.Price + c.motherboard.Price
}

type Builder struct {
	computer *Computer
}

func NewBuilder() *Builder {
	return &Builder{
		computer: &Computer{
			cpu: CPU{
				Component: Component{
					Name:  "cpu default",
					Price: 50.05,
				},
			},
			memory: Memory{
				Component: Component{
					Name:  "memory default",
					Price: 50.05,
				},
			},
			storage: Storage{
				Component: Component{
					Name:  "storage default",
					Price: 20.05,
				},
			},
			motherboard: Motherboard{
				Component: Component{
					Name:  "motherboard default",
					Price: 35.05,
				},
			},
		},
	}
}

func (b *Builder) WithCPU(cpu CPU) *Builder {
	b.computer.cpu = cpu
	return b
}

func (b *Builder) WithMemory(memory Memory) *Builder {
	b.computer.memory = memory
	return b
}
func (b *Builder) WithStorage(storage Storage) *Builder {
	b.computer.storage = storage
	return b
}
func (b *Builder) WithMotherboard(motherboard Motherboard) *Builder {
	b.computer.motherboard = motherboard
	return b
}

func (b *Builder) Build() *Computer {
	return b.computer
}
