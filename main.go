package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	du "internal/du"
	tui "internal/tui"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

/*
	This var sets up the root command and then all other commands. The root command, according to Cobra's structure, is the first thing we hit when we run the program.
	Imagine it as an automatic constructor that's allowing us to run an instance of this program.
*/
var rootCmd = &cobra.Command{
	Use: "put usage example here",
	//TraverseChildren: true,
	Short: "This program shows disk usage",
	Long:  "Put longer version of Short here",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print(args)
		v, _ := cmd.Flags().GetBool("version")
		if v {
			version()
		} else {
			//log.Fatal("version unavailable")
			fmt.Println("version not checked")
		}
		l := logFlag
		if l != "" {
			logFile, err = os.OpenFile(l, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
			if err != nil {
				err = fmt.Errorf("error opening log file: %w", err)
			}
		}

		o := outputFlag
		if o == "-" {
			outputFile = os.Stdout
		} else {
			outputFile, err = os.OpenFile(o, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
			if err != nil {
				err = fmt.Errorf("error setting output file: %w", err)
			}
		}

		defer func() {
			cerr := logFile.Close()
			if err == nil {
				err = cerr
			}
		}()
		log.SetOutput(logFile)

		if len(args) == 1 {
			dir, _ = filepath.Abs(args[0])
		} else {
			dir = "."
		}
		f := inputFile
		if f != "" {
			//needs to have logic for setting input file
			fmt.Println(f)
		}
		e := extendedFlag
		if e {
			fmt.Println("Needs implementation of extended information output")
		}
		ic := icFlag
		if ic {
			fmt.Println("Needs implementation of ignore configuration")
		}
		x := xFlag
		if x {
			fmt.Println("Needs implementation of filesystem boundaries")
		}
		c := cfsFlag
		if c {
			fmt.Println("Needs implementation of filesystem boundaries")
		}
		ex := exclude
		if len(ex) != 0 {
			for i := 0; i < len(ex); i++ {
				fmt.Println(ex[i])
			}
		}
		bX := bigXFlag
		if bX != "" {
			fmt.Println("Needs implementation of Exclusion Files")
		}
		sym := symLinkFlag
		if sym {
			fmt.Println("Needs implementation of symLink following")
		}
	},
}

var (
	//Scan and mode selection options
	symLinkFlag  bool
	bigXFlag     string
	exclude      []string
	cfsFlag      bool
	xFlag        bool
	icFlag       bool
	extendedFlag bool
	versionFlag  bool
	inputFile    string
	outputFlag   string
	outputFile   io.Writer
	logFlag      string
	logFile      *os.File
	err          error
	dir          string
	//interface options
	zeroFlag     bool
	oneFlag      bool
	twoFlag      bool
	qFlag        bool
	fastFlag     bool
	slowFlag     bool
	eShellFlag   bool
	dShellFlag   bool
	eDeleteFlag  bool
	dDeleteFlag  bool
	eRefreshFlag bool
	dRefreshFlag bool
	rFlag        bool
	siFlag       bool
	duFlag       bool
	shFlag       bool
	sicFlag      bool
	smtFlag      bool
	sgFlag       bool
	spFlag       bool
	gStyleFlag   string
	sColumnFlag  string
	sortFlag     string
	gdFlag       bool
	cqFlag       bool
	cdFlag       bool
	colorFlag    string
)

func version() {
	fmt.Println("Version goes here")
}

func init() {
	flags := rootCmd.Flags()

	//Scan and mode selection option flags
	flags.StringVarP(&outputFlag, "output-file", "o", "", "-o [FILE] defines file for data output")
	flags.StringVarP(&inputFile, "input-file", "f", "", "-f [FILE] defines file for data input")
	flags.BoolVarP(&versionFlag, "version", "v", false, "-v shows the current version of godu")
	flags.BoolVarP(&extendedFlag, "extended", "e", false, "-e enables extended information mode")
	flags.BoolVar(&icFlag, "ignore-config", false, "--ignore-config prevents godu from attempting to load any configuration files")
	flags.BoolVarP(&xFlag, "one-file-system", "x", false, "-x prevents godu from crossing filesystem boundaries, i.e. only count files and directories on the same filesystem as the directory being scanned")
	flags.BoolVar(&cfsFlag, "cross-file-system", false, "--cross-file-system allows godu to cross filesystem boundaries. This is the default, but can be specified to overrule a previously given '-x'")
	flags.StringArrayVar(&exclude, "exclude", exclude, "--exclude [PATTERN] excludes files that match PATTERN. The files will still be displayed by default, but are not counted towards the disk usage statistics. This argument can be added multiple times to add more patterns.")
	flags.StringVarP(&bigXFlag, "exclude-from", "X", "", "-X [FILE], --exclude-from [FILE] Exclude files that match any pattern in FILE. Patterns should be separated by a newline.")
	flags.BoolVarP(&symLinkFlag, "follow-symlinks", "L", false, "-L, --follow-symlinks follows symlinks and counts the size of the file they point to.")

	//interface option flags
	flags.BoolVar(&zeroFlag, "0", true, "Don't give any feedback while scanning a directory or importing a file, other than when a fatal error occurs. This option is the default when exporting to standard output.")
	flags.BoolVar(&oneFlag, "1", false, "Similar to -0, but does give feedback on the scanning progress with a single line of output. This option is the default when exporting to a file.")
	flags.BoolVar(&twoFlag, "2", false, "Provide a full-screen ncurses interface while scanning a directory or importing a file. This is the only interface that provides feedback on any non-fatal errors while scanning.")
	flags.BoolVar(&qFlag, "q", false, "Change the UI update interval while scanning or importing. This can be decreased to once every 2 seconds with -q or --slow-ui-updates. This feature can be used to save bandwidth over remote connections, but has no effect when -0 is used.")
	flags.BoolVar(&fastFlag, "fast-ui-updates", false, "Change the UI update interval while scanning or importing to 10 times per second. This option has no effect when -0 is used.")
	flags.BoolVar(&slowFlag, "slow-ui-updates", false, "Change the UI update interval while scanning or importing. This can be decreased to once every 2 seconds with -q or --slow-ui-updates. This feature can be used to save bandwidth over remote connections, but has no effect when -0 is used.")
	flags.BoolVar(&eShellFlag, "enable-shell", true, "Enable shell spawning from the browser. This feature is enabled by default when scanning a live directory and disabled when importing from file.")
	flags.BoolVar(&dShellFlag, "disable-shell", false, "Disable shell spawning from the browser. This feature is enabled by default when scanning a live directory and disabled when importing from file.")
	flags.BoolVar(&eDeleteFlag, "enable-delete", true, "Enable the built-in file deletion feature. This feature is enabled by default when scanning a live directory and disabled when importing from file. Explicitly disabling the deletion feature can work as a safeguard to prevent accidental data loss.")
	flags.BoolVar(&dDeleteFlag, "disable-delete", false, "Disable the built-in file deletion feature. This feature is enabled by default when scanning a live directory and disabled when importing from file. Explicitly disabling the deletion feature can work as a safeguard to prevent accidental data loss.")
	flags.BoolVar(&eRefreshFlag, "enable-refresh", true, "Enable directory refreshing from the browser. This feature is enabled by default when scanning a live directory and disabled when importing from file.")
	flags.BoolVar(&dRefreshFlag, "disable-refresh", false, "Disable directory refreshing from the browser. This feature is enabled by default when scanning a live directory and disabled when importing from file.")
	flags.BoolVar(&rFlag, "r", false, "Read-only mode. When given once, this is an alias for --disable-delete, when given twice it will also add --disable-shell, thus ensuring that there is no way to modify the file system from within godu.")
	flags.BoolVar(&siFlag, "si", false, "List sizes using base 10 prefixes, that is, powers of 1000 (KB, MB, etc), as defined in the International System of Units (SI), instead of the usual base 2 prefixes, that is, powers of 1024 (KiB, MiB, etc).")
	flags.BoolVar(&duFlag, "disk-usage", false, "Select whether to display disk usage (default) or apparent sizes. Can also be toggled in the browser with the 'a' key.")
	flags.BoolVar(&shFlag, "show-hidden", true, "Show (default) or hide 'hidden' and excluded files. Can also be toggled in the browser with the 'e' key.")
	flags.BoolVar(&sicFlag, "show-itemcount", false, "Show or hide (default) the item counts column. Can also be toggled in the browser with the 'c' key.")
	flags.BoolVar(&smtFlag, "show-mtime", false, "Show or hide (default) the last modification time column. Can also be toggled in the browser with the 'm' key. This option is ignored when not in extended mode (see -e).")
	flags.BoolVar(&sgFlag, "show-graph", true, "Show (default) or hide the relative size bar column. Can also be toggled in the browser with the 'g' key.")
	flags.BoolVar(&spFlag, "show-percent", true, "Show (default) or hide the relative size percent column. Can also be toggled in the browser with the 'g' key.")
	flags.StringVar(&gStyleFlag, "graph-style", "", "graph-style [OPTION]: Change the way that the relative size bar column is drawn. Recognized values are hash to draw ASCII # characters (default and most portable), half-block to use half-block drawing characters or eighth-block to use eighth-block drawing characters. Eighth-block characters are the most precise but may not render correctly in all terminals.")
	flags.StringVar(&sColumnFlag, "shared-column", "shared", "shared-column [OPTION]: Set to off to disable the shared size column for directories, shared (default) to display shared directory sizes as a separate column or unique to display unique directory sizes as a separate column. These options can also be cycled through in the browser with the 'u' key.")
	flags.StringVar(&sortFlag, "sort", "disk-usage", "sort [COLUMN]: Change the default column to sort on. Accepted values are disk-usage (the default), name, apparent-size, itemcount or mtime. The latter only makes sense in extended mode, see -e. The column can be suffixed with -asc or -desc to set the order to ascending or descending, respectively. e.g. --sort=name-desc will sort by name in descending order.")
	flags.BoolVar(&gdFlag, "group-directories-first", false, "Sort (or not) directories before files.")
	flags.BoolVar(&cqFlag, "confirm-quit", true, "Require a confirmation before quitting ncdu. Very helpful when you accidentally press 'q' during or after a very long scan.")
	flags.BoolVar(&cdFlag, "confirm-delete", true, "Require a confirmation before deleting a file or directory. Enabled by default, but can be disabled if you're absolutely sure you won't accidentally press 'd'.")
	flags.StringVar(&colorFlag, "color", "", "color [SCHEME]: Select a color scheme. The following schemes are recognized: off to disable colors, dark for a color scheme intended for dark backgrounds and dark-bg for a variation of the dark color scheme that also works in terminals with a light background. The default is dark-bg unless the NO_COLOR environment variable is set.")

}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}

	directory := "."
	hidden := false
	defaultOrdering := "name"
	directoryOrder := true
	diskUsage := true
	// percentage := true, graph, both, none
	uniqCol := false
	modifyTime := false

	files, err := du.ListFilesRecursivelyInParallel(directory)
	if err != nil {
		log.Fatalln(err)
	}

	initialModel := tui.Model{
		CurrentDirectory: directory,
		ShowHidden:       hidden,
		Order:            defaultOrdering,
		DirectoryFirst:   directoryOrder,
		ShowDiskUsage:    diskUsage,
		ShowUniqCol:      uniqCol,
		ModifyTime:       modifyTime,
		Files:            files,
	}

	p := tea.NewProgram(tui.NewModel(initialModel), tea.WithAltScreen())
	if err := p.Start(); err != nil {
		log.Fatal(err)
	}
}
