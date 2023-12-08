package build

import (
	"os"
	"os/exec"
	"testing"

	go_fuzz_utils "github.com/trailofbits/go-fuzz-utils"
)

func GetTypeProvider(data []byte) (*go_fuzz_utils.TypeProvider, error) {
	tp, err := go_fuzz_utils.NewTypeProvider(data)
	if err != nil {
		return nil, err
	}
	err = tp.SetParamsStringBounds(0, 1024)
	if err != nil {
		return nil, err
	}
	err = tp.SetParamsSliceBounds(0, 4096)
	if err != nil {
		return nil, err
	}
	err = tp.SetParamsBiases(0, 0, 0, 0)
	if err != nil {
		return nil, err
	}
	return tp, nil
}

func Fuzz_Nosy_ChecksumDB_DownloadFile__(f *testing.F) {
	f.Fuzz(func(t *testing.T, file string, url string, dstPath string) {
		db := MustLoadChecksums(file)
		db.DownloadFile(url, dstPath)
	})
}

func Fuzz_Nosy_ChecksumDB_Verify__(f *testing.F) {
	f.Fuzz(func(t *testing.T, file string, path string) {
		db := MustLoadChecksums(file)
		db.Verify(path)
	})
}

func Fuzz_Nosy_ChecksumDB_findHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, file string, basename string, hash string) {
		db := MustLoadChecksums(file)
		db.findHash(basename, hash)
	})
}

func Fuzz_Nosy_GoToolchain_Go__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var g *GoToolchain
		fill_err = tp.Fill(&g)
		if fill_err != nil {
			return
		}
		var command string
		fill_err = tp.Fill(&command)
		if fill_err != nil {
			return
		}
		var args []string
		fill_err = tp.Fill(&args)
		if fill_err != nil {
			return
		}
		if g == nil {
			return
		}

		g.Go(command, args...)
	})
}

func Fuzz_Nosy_GoToolchain_goTool__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var g *GoToolchain
		fill_err = tp.Fill(&g)
		if fill_err != nil {
			return
		}
		var command string
		fill_err = tp.Fill(&command)
		if fill_err != nil {
			return
		}
		var args []string
		fill_err = tp.Fill(&args)
		if fill_err != nil {
			return
		}
		if g == nil {
			return
		}

		g.goTool(command, args...)
	})
}

func Fuzz_Nosy_TarballArchive_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var a *TarballArchive
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		if a == nil {
			return
		}

		a.Close()
	})
}

func Fuzz_Nosy_TarballArchive_Directory__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var a *TarballArchive
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		var name string
		fill_err = tp.Fill(&name)
		if fill_err != nil {
			return
		}
		if a == nil {
			return
		}

		a.Directory(name)
	})
}

// skipping Fuzz_Nosy_TarballArchive_Header__ because parameters include func, chan, or unsupported interface: io/fs.FileInfo

func Fuzz_Nosy_ZipArchive_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var a *ZipArchive
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		if a == nil {
			return
		}

		a.Close()
	})
}

func Fuzz_Nosy_ZipArchive_Directory__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var a *ZipArchive
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		var name string
		fill_err = tp.Fill(&name)
		if fill_err != nil {
			return
		}
		if a == nil {
			return
		}

		a.Directory(name)
	})
}

// skipping Fuzz_Nosy_ZipArchive_Header__ because parameters include func, chan, or unsupported interface: io/fs.FileInfo

func Fuzz_Nosy_downloadWriter_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var dst *os.File
		fill_err = tp.Fill(&dst)
		if fill_err != nil {
			return
		}
		var size int64
		fill_err = tp.Fill(&size)
		if fill_err != nil {
			return
		}
		if dst == nil {
			return
		}

		w := newDownloadWriter(dst, size)
		w.Close()
	})
}

func Fuzz_Nosy_downloadWriter_Write__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var dst *os.File
		fill_err = tp.Fill(&dst)
		if fill_err != nil {
			return
		}
		var size int64
		fill_err = tp.Fill(&size)
		if fill_err != nil {
			return
		}
		var buf []byte
		fill_err = tp.Fill(&buf)
		if fill_err != nil {
			return
		}
		if dst == nil {
			return
		}

		w := newDownloadWriter(dst, size)
		w.Write(buf)
	})
}

// skipping Fuzz_Nosy_Archive_Close__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/internal/build.Archive

// skipping Fuzz_Nosy_Archive_Directory__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/internal/build.Archive

// skipping Fuzz_Nosy_Archive_Header__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/internal/build.Archive

func Fuzz_Nosy_AzureBlobstoreList__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var config AzureBlobstoreConfig
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}

		AzureBlobstoreList(config)
	})
}

func Fuzz_Nosy_DownloadAndVerifyChecksums__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var csdb *ChecksumDB
		fill_err = tp.Fill(&csdb)
		if fill_err != nil {
			return
		}
		if csdb == nil {
			return
		}

		DownloadAndVerifyChecksums(csdb)
	})
}

func Fuzz_Nosy_DownloadGo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var csdb *ChecksumDB
		fill_err = tp.Fill(&csdb)
		if fill_err != nil {
			return
		}
		if csdb == nil {
			return
		}

		DownloadGo(csdb)
	})
}

func Fuzz_Nosy_FindMainPackages__(f *testing.F) {
	f.Fuzz(func(t *testing.T, dir string) {
		FindMainPackages(dir)
	})
}

func Fuzz_Nosy_MustRun__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cmd *exec.Cmd
		fill_err = tp.Fill(&cmd)
		if fill_err != nil {
			return
		}
		if cmd == nil {
			return
		}

		MustRun(cmd)
	})
}

func Fuzz_Nosy_MustRunCommand__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cmd string
		fill_err = tp.Fill(&cmd)
		if fill_err != nil {
			return
		}
		var args []string
		fill_err = tp.Fill(&args)
		if fill_err != nil {
			return
		}

		MustRunCommand(cmd, args...)
	})
}

func Fuzz_Nosy_NewArchive__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var file *os.File
		fill_err = tp.Fill(&file)
		if fill_err != nil {
			return
		}
		if file == nil {
			return
		}

		NewArchive(file)
	})
}

func Fuzz_Nosy_PGPKeyID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, pgpkey string) {
		PGPKeyID(pgpkey)
	})
}

// skipping Fuzz_Nosy_Render__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_RenderString__ because parameters include func, chan, or unsupported interface: interface{}

func Fuzz_Nosy_RunGit__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var args []string
		fill_err = tp.Fill(&args)
		if fill_err != nil {
			return
		}

		RunGit(args...)
	})
}

func Fuzz_Nosy_Version__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var csdb *ChecksumDB
		fill_err = tp.Fill(&csdb)
		if fill_err != nil {
			return
		}
		var version string
		fill_err = tp.Fill(&version)
		if fill_err != nil {
			return
		}
		if csdb == nil {
			return
		}

		Version(csdb, version)
	})
}

func Fuzz_Nosy_firstLine__(f *testing.F) {
	f.Fuzz(func(t *testing.T, s string) {
		firstLine(s)
	})
}

func Fuzz_Nosy_getDate__(f *testing.F) {
	f.Fuzz(func(t *testing.T, commit string) {
		getDate(commit)
	})
}

func Fuzz_Nosy_printArgs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var args []string
		fill_err = tp.Fill(&args)
		if fill_err != nil {
			return
		}

		printArgs(args)
	})
}

func Fuzz_Nosy_readGitFile__(f *testing.F) {
	f.Fuzz(func(t *testing.T, file string) {
		readGitFile(file)
	})
}

// skipping Fuzz_Nosy_render__ because parameters include func, chan, or unsupported interface: interface{}
