package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
)

func usage(exit int) {
	fmt.Printf("Usage: atws <layer|workspace|move-layer|move-workspace|update|reset> <arg>\n")
	os.Exit(exit)
}

func main() {
	if len(os.Args) == 2 {
		switch os.Args[1] {
		case "update":
			cfg, err := GetConfig()
			panicErr("GetConfig", err)

			update(cfg)

		case "reset":
			cfg, err := GetConfig()
			panicErr("GetConfig", err)

			cfg.Layer = 0
			cfg.Workspace = 0

			panicErr("SetConfig", SetConfig(cfg))

		default:
			usage(1)
		}

		os.Exit(0)
	}

	if len(os.Args) != 3 {
		usage(1)
	}

	cfg, err := GetConfig()
	panicErr("GetConfig", err)

	i, err := strconv.Atoi(os.Args[2])
	panicErr("Atoi", err)

	switch os.Args[1] {
	case "layer":
		cfg.Layer = i

	case "workspace":
		cfg.Workspace = i

	case "move-workspace":
		cfg.Workspace = i

		move(cfg)
		os.Exit(0)

	case "move-layer":
		cfg.Layer = i

		move(cfg)
		os.Exit(0)

	default:
		usage(1)
	}

	panicErr("SetConfig", SetConfig(cfg))
	update(cfg)
}

func panicErr(f string, err error) {
	if err != nil {
		fmt.Printf("%s() -> %s\n", f, err)
		os.Exit(1)
	}
}

func update(cfg *Config) {
	cmd := exec.Command("i3-msg", "workspace",
		fmt.Sprintf("%d", cfg.Workspace+(cfg.MaxLayer*cfg.Layer)),
	)
	panicErr("exec.Run", cmd.Run())
}

func move(cfg *Config) {
	cmd := exec.Command("i3-msg", "move", "container", "to", "workspace", "number",
		fmt.Sprintf("%d", cfg.Workspace+(cfg.MaxLayer*cfg.Layer)),
	)
	panicErr("exec.Run", cmd.Run())
}
