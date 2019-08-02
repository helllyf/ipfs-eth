package commands

//ganache-cli -m "much repair shock carbon improve miss forget sock include bullet interest solution"
import (
	"errors"
	"fmt"
	"io"
	"os"
	"path"
	"strconv"
	"strings"

	cmds "github.com/ipfs/go-ipfs-cmds"
	files "github.com/ipfs/go-ipfs-files"
	"github.com/ipfs/go-ipfs/core/commands/cmdenv"
	"github.com/ipfs/go-ipfs/ethbridge"
	coreiface "github.com/ipfs/interface-go-ipfs-core"
	"github.com/ipfs/interface-go-ipfs-core/options"
	mh "github.com/multiformats/go-multihash"
	"gopkg.in/cheggaaa/pb.v1"
)

const groupName = "group-name"

type TellOutput struct {
	test string
}

func ethClientInit(env cmds.Environment) {
	rootPath,err := cmdenv.GetConfigRoot(env)
	if err != nil {
		log.Fatal(err)
	}
	ethbridge.InitEthClient(rootPath)
}



var TellEthCmd = &cmds.Command{
	Helptext: cmds.HelpText{
		Tagline:          "Get ethereum.",
		ShortDescription: "Returns something from eth.",
	},
	Subcommands: map[string]*cmds.Command{
		"tell": 			Tell,
		"showGroups":		ShowGroups,
		"addGroup":			AddGroup,
		"addUserToGroup":	AddUserToGroup,
		"setPermission" :	SetPermission,
		"ls":				EthLsCmd,
	},
}

// 执行顺序 prerun postrun run ...
var Tell = &cmds.Command{
	Helptext: cmds.HelpText{
		Tagline:          "Get ethereum.",
		ShortDescription: "Returns something from eth.",
	},
	Run: func(req *cmds.Request, res cmds.ResponseEmitter, env cmds.Environment) error {
		ethClientInit(env)
		balance := ethbridge.GetBalance()
		fmt.Println(balance)
		ethbridge.TestRouter()
		return cmds.EmitOnce(res, &TellOutput{
			test: "file system",
		})
	},
	Encoders: cmds.EncoderMap{
		cmds.Text: cmds.MakeTypedEncoder(func(req *cmds.Request, w io.Writer, out *TellOutput) error {
			return nil
		}),
	},
	Type: TellOutput{},
}

var ShowGroups = &cmds.Command{
	Helptext: cmds.HelpText{
		Tagline:			"show user Groups",
		ShortDescription:	"show name and address of User's Groups",
	},
	PreRun: func(req *cmds.Request, env cmds.Environment) error {
		ethClientInit(env)
		return nil
	},
	Run: func(request *cmds.Request, emitter cmds.ResponseEmitter, environment cmds.Environment) error {
		ethbridge.ShowUserGroups(true)
		return cmds.EmitOnce(emitter, &TellOutput{
			test: "file system",
		})
	},
	Encoders: cmds.EncoderMap{
		cmds.Text: cmds.MakeTypedEncoder(func(req *cmds.Request,w io.Writer,out *TellOutput) error {
			return nil
		}),
	},
}

var AddGroup = &cmds.Command{
	Helptext: cmds.HelpText{
		Tagline:			"add Groups",
		ShortDescription:	"add name of User's Groups,solidity will create new group contracts",
	},
	Arguments: []cmds.Argument{
		cmds.StringArg("groupname", true, true, "The name to a group to be added to groups.").EnableStdin(),
	},
	PreRun: func(req *cmds.Request, env cmds.Environment) error {
		ethClientInit(env)
		return nil
	},
	Run: func(request *cmds.Request, emitter cmds.ResponseEmitter, environment cmds.Environment) error {
		_name := request.Arguments[0]
		ethbridge.AddGroup(_name)
		return cmds.EmitOnce(emitter, &TellOutput{
			test: "file system",
		})
	},
	Encoders: cmds.EncoderMap{
		cmds.Text: cmds.MakeTypedEncoder(func(req *cmds.Request,w io.Writer,out *TellOutput) error {
			return nil
		}),
	},
}

var AddUserToGroup = &cmds.Command{
	Helptext: cmds.HelpText{
		Tagline:			"add User to Groups",
		ShortDescription:	"add name of User's Groups,solidity will create new group contracts",
	},
	Arguments: []cmds.Argument{
		cmds.StringArg("userAddress", true, true, "user address add to group.").EnableStdin(),
		cmds.StringArg("groupName", true, true, "group name.").EnableStdin(),
		cmds.StringArg("permission", true, true, "permission between 1 to 7.").EnableStdin(),
	},
	PreRun: func(req *cmds.Request, env cmds.Environment) error {
		ethClientInit(env)
		return nil
	},
	Run: func(request *cmds.Request, emitter cmds.ResponseEmitter, environment cmds.Environment) error {
		_userAddr 	:= request.Arguments[0]
		_groupName	:= request.Arguments[1]
		_p,err := strconv.Atoi(request.Arguments[2])
		if err != nil {
			log.Fatal(err)
		}
		ethbridge.AddUser(_userAddr, _groupName, uint(_p))
		return cmds.EmitOnce(emitter, &TellOutput{
			test: "file system",
		})
	},
	Encoders: cmds.EncoderMap{
		cmds.Text: cmds.MakeTypedEncoder(func(req *cmds.Request,w io.Writer,out *TellOutput) error {
			return nil
		}),
	},
}

var SetPermission =  &cmds.Command{
	Helptext: cmds.HelpText{
		Tagline:			"add permission to Dic",
		ShortDescription:	"",
	},
	Arguments: []cmds.Argument{
		cmds.FileArg("path", true, true, "The path to a file to be added to ipfs.").EnableRecursive().EnableStdin(),
	},
	Options: []cmds.Option{
		cmds.OptionRecursivePath, // a builtin option that allows recursive paths (-r, --recursive)
		cmds.OptionDerefArgs,     // a builtin option that resolves passed in filesystem links (--dereference-args)
		cmds.OptionStdinName,     // a builtin option that optionally allows wrapping stdin into a named file
		cmds.OptionHidden,
		cmds.BoolOption(quietOptionName, "q", "Write minimal output."),
		cmds.BoolOption(quieterOptionName, "Q", "Write only final hash."),
		cmds.BoolOption(silentOptionName, "Write no output."),
		cmds.BoolOption(progressOptionName, "p", "Stream progress data."),
		cmds.BoolOption(trickleOptionName, "t", "Use trickle-dag format for dag generation."),
		cmds.BoolOption(onlyHashOptionName, "n", "Only chunk and hash - do not write to disk."),
		cmds.BoolOption(wrapOptionName, "w", "Wrap files with a directory object."),
		cmds.StringOption(chunkerOptionName, "s", "Chunking algorithm, size-[bytes] or rabin-[min]-[avg]-[max]").WithDefault("size-262144"),
		cmds.BoolOption(pinOptionName, "Pin this object when adding.").WithDefault(true),
		cmds.BoolOption(rawLeavesOptionName, "Use raw blocks for leaf nodes. (experimental)"),
		cmds.BoolOption(noCopyOptionName, "Add the file using filestore. Implies raw-leaves. (experimental)"),
		cmds.BoolOption(fstoreCacheOptionName, "Check the filestore for pre-existing blocks. (experimental)"),
		cmds.IntOption(cidVersionOptionName, "CID version. Defaults to 0 unless an option that depends on CIDv1 is passed. (experimental)"),
		cmds.StringOption(hashOptionName, "Hash function to use. Implies CIDv1 if not sha2-256. (experimental)").WithDefault("sha2-256"),
		cmds.BoolOption(inlineOptionName, "Inline small blocks into CIDs. (experimental)"),
		cmds.IntOption(inlineLimitOptionName, "Maximum block size to inline. (experimental)").WithDefault(32),

		cmds.StringOption(groupName,"g","Add group name for dictionary"),
	},
	PreRun: func(req *cmds.Request, env cmds.Environment) error {
		ethClientInit(env)

		quiet, _ := req.Options[quietOptionName].(bool)
		quieter, _ := req.Options[quieterOptionName].(bool)
		quiet = quiet || quieter

		silent, _ := req.Options[silentOptionName].(bool)

		if quiet || silent {
			return nil
		}

		// ipfs cli progress bar defaults to true unless quiet or silent is used
		_, found := req.Options[progressOptionName].(bool)
		if !found {
			req.Options[progressOptionName] = true
		}

		return nil
	},
	Run: func(req *cmds.Request, res cmds.ResponseEmitter, env cmds.Environment) error {
		api, err := cmdenv.GetApi(env, req)
		if err != nil {
			return err
		}

		progress, _ := req.Options[progressOptionName].(bool)
		trickle, _ := req.Options[trickleOptionName].(bool)
		wrap, _ := req.Options[wrapOptionName].(bool)
		hash, _ := req.Options[onlyHashOptionName].(bool)
		silent, _ := req.Options[silentOptionName].(bool)
		chunker, _ := req.Options[chunkerOptionName].(string)
		dopin, _ := req.Options[pinOptionName].(bool)
		rawblks, rbset := req.Options[rawLeavesOptionName].(bool)
		nocopy, _ := req.Options[noCopyOptionName].(bool)
		fscache, _ := req.Options[fstoreCacheOptionName].(bool)
		cidVer, cidVerSet := req.Options[cidVersionOptionName].(int)
		hashFunStr, _ := req.Options[hashOptionName].(string)
		inline, _ := req.Options[inlineOptionName].(bool)
		inlineLimit, _ := req.Options[inlineLimitOptionName].(int)

		hashFunCode, ok := mh.Names[strings.ToLower(hashFunStr)]
		if !ok {
			return fmt.Errorf("unrecognized hash function: %s", strings.ToLower(hashFunStr))
		}


		enc, err := cmdenv.GetCidEncoder(req)
		if err != nil {
			return err
		}

		toadd := req.Files
		//给文件包个目录
		if wrap {
			toadd = files.NewSliceDirectory([]files.DirEntry{
				files.FileEntry("", req.Files),
			})
		}


		opts := []options.UnixfsAddOption{
			options.Unixfs.Hash(hashFunCode),

			options.Unixfs.Inline(inline),
			options.Unixfs.InlineLimit(inlineLimit),

			options.Unixfs.Chunker(chunker),

			options.Unixfs.Pin(dopin),
			options.Unixfs.HashOnly(hash),
			options.Unixfs.FsCache(fscache),
			options.Unixfs.Nocopy(nocopy),

			options.Unixfs.Progress(progress),
			options.Unixfs.Silent(silent),
		}

		if cidVerSet {
			opts = append(opts, options.Unixfs.CidVersion(cidVer))
		}

		if rbset {
			opts = append(opts, options.Unixfs.RawLeaves(rawblks))
		}

		if trickle {
			opts = append(opts, options.Unixfs.Layout(options.TrickleLayout))
		}

		opts = append(opts, nil) // events option placeholder

		var added int
		addit := toadd.Entries()
		for addit.Next() {
			_, dir := addit.Node().(files.Directory) //类型断言，判断node是否为该接口类型
			errCh := make(chan error, 1)
			events := make(chan interface{}, adderOutChanSize)
			opts[len(opts)-1] = options.Unixfs.Events(events)

			go func() {
				var err error
				defer close(events)
				//Unixfs returns the UnixfsAPI interface implementation backed by the go-ipfs node
				_, err = api.Unixfs().Add(req.Context, addit.Node(), opts...)
				errCh <- err
			}()

			for event := range events {
				output, ok := event.(*coreiface.AddEvent)
				if !ok {
					return errors.New("unknown event type")
				}

				h := ""
				if output.Path != nil {
					h = enc.Encode(output.Path.Cid())
				}

		//		_filePath,ok := addit.Node().(files.FileInfo)
		//		filePath := _filePath.AbsPath() + output.Name
		//		fmt.Printf("%s is file: %v\n", filePath, isFile(filePath))

				if !dir && addit.Name() != "" {
					output.Name = addit.Name()
				} else {
					output.Name = path.Join(addit.Name(), output.Name)
				}

				if err := res.Emit(&AddEvent{
					Name:  output.Name,
					Hash:  h,
					Bytes: output.Bytes,
					Size:  output.Size,
				}); err != nil {
					return err
				}
			}

			if err := <-errCh; err != nil {
				return err
			}
			added++
		}

		if addit.Err() != nil {
			return addit.Err()
		}

		if added == 0 {
			return fmt.Errorf("expected a file argument")
		}

		return nil
	},
	PostRun: cmds.PostRunMap{
		cmds.CLI: func(res cmds.Response, re cmds.ResponseEmitter) error {

			sizeChan := make(chan int64, 1)
			outChan := make(chan interface{})
			req := res.Request()

			// Could be slow.
			go func() {
				size, err := req.Files.Size()
				if err != nil {
					log.Warningf("error getting files size: %s", err)
					// see comment above
					return
				}

				sizeChan <- size
			}()

			progressBar := func(wait chan struct{}) {
				defer close(wait)

				quiet, _ := req.Options[quietOptionName].(bool)
				quieter, _ := req.Options[quieterOptionName].(bool)
				quiet = quiet || quieter

				progress, _ := req.Options[progressOptionName].(bool)

				var bar *pb.ProgressBar
				if progress {
					bar = pb.New64(0).SetUnits(pb.U_BYTES)
					bar.ManualUpdate = true
					bar.ShowTimeLeft = false
					bar.ShowPercent = false
					bar.Output = os.Stderr
					bar.Start()
				}

				lastFile := ""
				lastHash := ""
				var totalProgress, prevFiles, lastBytes int64

			LOOP:
				for {
					select {
					case out, ok := <-outChan:
						if !ok {
							if quieter {
								fmt.Fprintln(os.Stdout, lastHash)
							}
							fmt.Println("/n")
							ethbridge.CommitToEth()
							break LOOP
						}
						output := out.(*AddEvent)
						if len(output.Hash) > 0 {
							lastHash = output.Hash
							if quieter {
								continue
							}

							if progress {
								// clear progress bar line before we print "added x" output
								fmt.Fprintf(os.Stderr, "\033[2K\r")
							}
							if quiet {
								fmt.Fprintf(os.Stdout, "%s\n", output.Hash)
							} else {
								fmt.Fprintf(os.Stdout, "added %s %s\n", output.Hash, output.Name)
								//设置权限
								ethbridge.SetPermission(output.Name,req.Options[groupName].(string),output.Hash)

							}

						} else {
							if !progress {
								continue
							}

							if len(lastFile) == 0 {
								lastFile = output.Name
							}
							if output.Name != lastFile || output.Bytes < lastBytes {
								prevFiles += lastBytes
								lastFile = output.Name
							}
							lastBytes = output.Bytes
							delta := prevFiles + lastBytes - totalProgress
							totalProgress = bar.Add64(delta)
						}

						if progress {
							bar.Update()
						}
					case size := <-sizeChan:
						if progress {
							bar.Total = size
							bar.ShowPercent = true
							bar.ShowBar = true
							bar.ShowTimeLeft = true
						}
					case <-req.Context.Done():

						// don't set or print error here, that happens in the goroutine below
						return
					}
				}

				if progress && bar.Total == 0 && bar.Get() != 0 {
					bar.Total = bar.Get()
					bar.ShowPercent = true
					bar.ShowBar = true
					bar.ShowTimeLeft = true
					bar.Update()
				}
			}

			if e := res.Error(); e != nil {
				close(outChan)
				return e
			}

			wait := make(chan struct{})
			go progressBar(wait)

			defer func() { <-wait }()
			defer close(outChan)
			for {
				v, err := res.Next()
				if err != nil {
					if err == io.EOF {
						return nil
					}

					return err
				}

				select {
				case outChan <- v:
				case <-req.Context.Done():
					return req.Context.Err()
				}
			}
		},
	},
	Type: AddEvent{},
}

var EthLsCmd = &cmds.Command{
	Helptext: cmds.HelpText{
		Tagline:          "Get ethereum.",
		ShortDescription: "Returns something from eth.",
	},
	Run: func(req *cmds.Request, res cmds.ResponseEmitter, env cmds.Environment) error {
		ethClientInit(env)
		ethbridge.PrintEthDic()
		return cmds.EmitOnce(res, &TellOutput{
			test: "file system",
		})
	},
	Encoders: cmds.EncoderMap{
		cmds.Text: cmds.MakeTypedEncoder(func(req *cmds.Request, w io.Writer, out *TellOutput) error {
			return nil
		}),
	},
	Type: TellOutput{},
}
