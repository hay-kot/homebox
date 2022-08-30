package automapper

type AutoMapperConf struct {
	OutDir string
}

func DefaultConf() *AutoMapperConf {
	return &AutoMapperConf{
		OutDir: "internal/mapper",
	}
}
