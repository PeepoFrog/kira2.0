package config

type Config struct {
	sekai  Sekai // Here we will place out config
	interx Interx
}

type Sekai struct {
	config ConfToml
}

type Interx struct {
}

type ConfToml struct {
}
