type Command struct {
    // Use is the one-line usage message.
    // Recommended syntax is as follows:
    //   [ ] identifies an optional argument. Arguments that are not enclosed in brackets are required.
    //   ... indicates that you can specify multiple values for the previous argument.
    //   |   indicates mutually exclusive information. You can use the argument to the left of the separator or the
    //       argument to the right of the separator. You cannot use both arguments in a single use of the command.
    //   { } delimits a set of mutually exclusive arguments when one of the arguments is required. If the arguments are
    //       optional, they are enclosed in brackets ([ ]).
    // Example: add [-F file | -D dir]... [-f format] profile
    Use string

    // Aliases is an array of aliases that can be used instead of the first word in Use.
    Aliases []string

    // SuggestFor is an array of command names for which this command will be suggested -
    // similar to aliases but only suggests.
    SuggestFor []string

    // Short is the short description shown in the 'help' output.
    Short string

    // The group id under which this subcommand is grouped in the 'help' output of its parent.
    GroupID string

    // Long is the long message shown in the 'help <this-command>' output.
    Long string

    // Example is examples of how to use the command.
    Example string

    // ValidArgs is list of all valid non-flag arguments that are accepted in shell completions
    ValidArgs []string
    // ValidArgsFunction is an optional function that provides valid non-flag arguments for shell completion.
    // It is a dynamic version of using ValidArgs.
    // Only one of ValidArgs and ValidArgsFunction can be used for a command.
    ValidArgsFunction func(cmd *Command, args []string, toComplete string) ([]string, ShellCompDirective)

    // Expected arguments
    Args PositionalArgs

    // ArgAliases is List of aliases for ValidArgs.
    // These are not suggested to the user in the shell completion,
    // but accepted if entered manually.
    ArgAliases []string

    // BashCompletionFunction is custom bash functions used by the legacy bash autocompletion generator.
    // For portability with other shells, it is recommended to instead use ValidArgsFunction
    BashCompletionFunction string

    // Deprecated defines, if this command is deprecated and should print this string when used.
    Deprecated string

    // Annotations are key/value pairs that can be used by applications to identify or
    // group commands.
    Annotations map[string]string

    // Version defines the version for this command. If this value is non-empty and the command does not
    // define a "version" flag, a "version" boolean flag will be added to the command and, if specified,
    // will print content of the "Version" variable. A shorthand "v" flag will also be added if the
    // command does not define one.
    Version string

    // The *Run functions are executed in the following order:
    //   * PersistentPreRun()
    //   * PreRun()
    //   * Run()
    //   * PostRun()
    //   * PersistentPostRun()
    // All functions get the same args, the arguments after the command name.
    //
    // PersistentPreRun: children of this command will inherit and execute.
    PersistentPreRun func(cmd *Command, args []string)
    // PersistentPreRunE: PersistentPreRun but returns an error.
    PersistentPreRunE func(cmd *Command, args []string) error
    // PreRun: children of this command will not inherit.
    PreRun func(cmd *Command, args []string)
    // PreRunE: PreRun but returns an error.
    PreRunE func(cmd *Command, args []string) error
    // Run: Typically the actual work function. Most commands will only implement this.
    Run func(cmd *Command, args []string)
    // RunE: Run but returns an error.
    RunE func(cmd *Command, args []string) error
    // PostRun: run after the Run command.
    PostRun func(cmd *Command, args []string)
    // PostRunE: PostRun but returns an error.
    PostRunE func(cmd *Command, args []string) error
    // PersistentPostRun: children of this command will inherit and execute after PostRun.
    PersistentPostRun func(cmd *Command, args []string)
    // PersistentPostRunE: PersistentPostRun but returns an error.
    PersistentPostRunE func(cmd *Command, args []string) error

    // groups for subcommands
    commandgroups []*Group

    // args is actual args parsed from flags.
    args []string
    // flagErrorBuf contains all error messages from pflag.
    flagErrorBuf *bytes.Buffer
    // flags is full set of flags.
    flags *flag.FlagSet
    // pflags contains persistent flags.
    pflags *flag.FlagSet
    // lflags contains local flags.
    lflags *flag.FlagSet
    // iflags contains inherited flags.
    iflags *flag.FlagSet
    // parentsPflags is all persistent flags of cmd's parents.
    parentsPflags *flag.FlagSet
    // globNormFunc is the global normalization function
    // that we can use on every pflag set and children commands
    globNormFunc func(f *flag.FlagSet, name string) flag.NormalizedName

    // usageFunc is usage func defined by user.
    usageFunc func(*Command) error
    // usageTemplate is usage template defined by user.
    usageTemplate string
    // flagErrorFunc is func defined by user and it's called when the parsing of
    // flags returns an error.
    flagErrorFunc func(*Command, error) error
    // helpTemplate is help template defined by user.
    helpTemplate string
    // helpFunc is help func defined by user.
    helpFunc func(*Command, []string)
    // helpCommand is command with usage 'help'. If it's not defined by user,
    // cobra uses default help command.
    helpCommand *Command
    // helpCommandGroupID is the group id for the helpCommand
    helpCommandGroupID string

    // completionCommandGroupID is the group id for the completion command
    completionCommandGroupID string

    // versionTemplate is the version template defined by user.
    versionTemplate string

    // inReader is a reader defined by the user that replaces stdin
    inReader io.Reader
    // outWriter is a writer defined by the user that replaces stdout
    outWriter io.Writer
    // errWriter is a writer defined by the user that replaces stderr
    errWriter io.Writer

    // FParseErrWhitelist flag parse errors to be ignored
    FParseErrWhitelist FParseErrWhitelist

    // CompletionOptions is a set of options to control the handling of shell completion
    CompletionOptions CompletionOptions

    // commandsAreSorted defines, if command slice are sorted or not.
    commandsAreSorted bool
    // commandCalledAs is the name or alias value used to call this command.
    commandCalledAs struct {
        name   string
        called bool
    }

    ctx context.Context

    // commands is the list of commands supported by this program.
    commands []*Command
    // parent is a parent command for this command.
    parent *Command
    // Max lengths of commands' string lengths for use in padding.
    commandsMaxUseLen         int
    commandsMaxCommandPathLen int
    commandsMaxNameLen        int

    // TraverseChildren parses flags on all parents before executing child command.
    TraverseChildren bool

    // Hidden defines, if this command is hidden and should NOT show up in the list of available commands.
    Hidden bool

    // SilenceErrors is an option to quiet errors down stream.
    SilenceErrors bool

    // SilenceUsage is an option to silence usage when an error occurs.
    SilenceUsage bool

    // DisableFlagParsing disables the flag parsing.
    // If this is true all flags will be passed to the command as arguments.
    DisableFlagParsing bool

    // DisableAutoGenTag defines, if gen tag ("Auto generated by spf13/cobra...")
    // will be printed by generating docs for this command.
    DisableAutoGenTag bool

    // DisableFlagsInUseLine will disable the addition of [flags] to the usage
    // line of a command when printing help or generating docs
    DisableFlagsInUseLine bool

    // DisableSuggestions disables the suggestions based on Levenshtein distance
    // that go along with 'unknown command' messages.
    DisableSuggestions bool

    // SuggestionsMinimumDistance defines minimum levenshtein distance to display suggestions.
    // Must be > 0.
    SuggestionsMinimumDistance int
}

func (*cobra.Command).AddCommand(cmds ...*cobra.Command)
func (*cobra.Command).AddGroup(groups ...*cobra.Group)
func (*cobra.Command).AllChildCommandsHaveGroup() bool
func (*cobra.Command).ArgsLenAtDash() int
func (*cobra.Command).CalledAs() string
func (*cobra.Command).CommandPath() string
func (*cobra.Command).CommandPathPadding() int
func (*cobra.Command).Commands() []*cobra.Command
func (*cobra.Command).ContainsGroup(groupID string) bool
func (*cobra.Command).Context() context.Context
func (*cobra.Command).DebugFlags()
func (*cobra.Command).ErrOrStderr() io.Writer
func (*cobra.Command).Execute() error
func (*cobra.Command).ExecuteC() (cmd *cobra.Command, err error)
func (*cobra.Command).ExecuteContext(ctx context.Context) error
func (*cobra.Command).ExecuteContextC(ctx context.Context) (*cobra.Command, error)
func (*cobra.Command).Find(args []string) (*cobra.Command, []string, error)
func (*cobra.Command).Flag(name string) (flag *pflag.Flag)
func (*cobra.Command).FlagErrorFunc() (f func(*cobra.Command, error) error)
func (*cobra.Command).Flags() *pflag.FlagSet
func (*cobra.Command).GenBashCompletion(w io.Writer) error
func (*cobra.Command).GenBashCompletionFile(filename string) error
func (*cobra.Command).GenBashCompletionFileV2(filename string, includeDesc bool) error
func (*cobra.Command).GenBashCompletionV2(w io.Writer, includeDesc bool) error
func (*cobra.Command).GenFishCompletion(w io.Writer, includeDesc bool) error
func (*cobra.Command).GenFishCompletionFile(filename string, includeDesc bool) error
func (*cobra.Command).GenPowerShellCompletion(w io.Writer) error
func (*cobra.Command).GenPowerShellCompletionFile(filename string) error
func (*cobra.Command).GenPowerShellCompletionFileWithDesc(filename string) error
func (*cobra.Command).GenPowerShellCompletionWithDesc(w io.Writer) error
func (*cobra.Command).GenZshCompletion(w io.Writer) error
func (*cobra.Command).GenZshCompletionFile(filename string) error
func (*cobra.Command).GenZshCompletionFileNoDesc(filename string) error
func (*cobra.Command).GenZshCompletionNoDesc(w io.Writer) error
func (*cobra.Command).GlobalNormalizationFunc() func(f *pflag.FlagSet, name string) pflag.NormalizedName
func (*cobra.Command).Groups() []*cobra.Group
func (*cobra.Command).HasAlias(s string) bool
func (*cobra.Command).HasAvailableFlags() bool
func (*cobra.Command).HasAvailableInheritedFlags() bool
func (*cobra.Command).HasAvailableLocalFlags() bool
func (*cobra.Command).HasAvailablePersistentFlags() bool
func (*cobra.Command).HasAvailableSubCommands() bool
func (*cobra.Command).HasExample() bool
func (*cobra.Command).HasFlags() bool
func (*cobra.Command).HasHelpSubCommands() bool
func (*cobra.Command).HasInheritedFlags() bool
func (*cobra.Command).HasLocalFlags() bool
func (*cobra.Command).HasParent() bool
func (*cobra.Command).HasPersistentFlags() bool
func (*cobra.Command).HasSubCommands() bool
func (*cobra.Command).Help() error
func (*cobra.Command).HelpFunc() func(*cobra.Command, []string)
func (*cobra.Command).HelpTemplate() string
func (*cobra.Command).InOrStdin() io.Reader
func (*cobra.Command).InheritedFlags() *pflag.FlagSet
func (*cobra.Command).InitDefaultCompletionCmd()
func (*cobra.Command).InitDefaultHelpCmd()
func (*cobra.Command).InitDefaultHelpFlag()
func (*cobra.Command).InitDefaultVersionFlag()
func (*cobra.Command).IsAdditionalHelpTopicCommand() bool
func (*cobra.Command).IsAvailableCommand() bool
func (*cobra.Command).LocalFlags() *pflag.FlagSet
func (*cobra.Command).LocalNonPersistentFlags() *pflag.FlagSet
func (*cobra.Command).MarkFlagCustom(name string, f string) error
func (*cobra.Command).MarkFlagDirname(name string) error
func (*cobra.Command).MarkFlagFilename(name string, extensions ...string) error
func (*cobra.Command).MarkFlagRequired(name string) error
func (*cobra.Command).MarkFlagsMutuallyExclusive(flagNames ...string)
func (*cobra.Command).MarkFlagsRequiredTogether(flagNames ...string)
func (*cobra.Command).MarkPersistentFlagDirname(name string) error
func (*cobra.Command).MarkPersistentFlagFilename(name string, extensions ...string) error
func (*cobra.Command).MarkPersistentFlagRequired(name string) error
func (*cobra.Command).MarkZshCompPositionalArgumentFile(argPosition int, patterns ...string) error
func (*cobra.Command).MarkZshCompPositionalArgumentWords(argPosition int, words ...string) error
func (*cobra.Command).Name() string
func (*cobra.Command).NameAndAliases() string
func (*cobra.Command).NamePadding() int
func (*cobra.Command).NonInheritedFlags() *pflag.FlagSet
func (*cobra.Command).OutOrStderr() io.Writer
func (*cobra.Command).OutOrStdout() io.Writer
func (*cobra.Command).Parent() *cobra.Command
func (*cobra.Command).ParseFlags(args []string) error
func (*cobra.Command).PersistentFlags() *pflag.FlagSet
func (*cobra.Command).Print(i ...interface{})
func (*cobra.Command).PrintErr(i ...interface{})
func (*cobra.Command).PrintErrf(format string, i ...interface{})
func (*cobra.Command).PrintErrln(i ...interface{})
func (*cobra.Command).Printf(format string, i ...interface{})
func (*cobra.Command).Println(i ...interface{})
func (*cobra.Command).RegisterFlagCompletionFunc(flagName string, f func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective)) error
func (*cobra.Command).RemoveCommand(cmds ...*cobra.Command)
func (*cobra.Command).ResetCommands()
func (*cobra.Command).ResetFlags()
func (*cobra.Command).Root() *cobra.Command
func (*cobra.Command).Runnable() bool
func (*cobra.Command).SetArgs(a []string)
func (*cobra.Command).SetCompletionCommandGroupID(groupID string)
func (*cobra.Command).SetContext(ctx context.Context)
func (*cobra.Command).SetErr(newErr io.Writer)
func (*cobra.Command).SetFlagErrorFunc(f func(*cobra.Command, error) error)
func (*cobra.Command).SetGlobalNormalizationFunc(n func(f *pflag.FlagSet, name string) pflag.NormalizedName)
func (*cobra.Command).SetHelpCommand(cmd *cobra.Command)
func (*cobra.Command).SetHelpCommandGroupID(groupID string)
func (*cobra.Command).SetHelpFunc(f func(*cobra.Command, []string))
func (*cobra.Command).SetHelpTemplate(s string)
func (*cobra.Command).SetIn(newIn io.Reader)
func (*cobra.Command).SetOut(newOut io.Writer)
func (*cobra.Command).SetOutput(output io.Writer)
func (*cobra.Command).SetUsageFunc(f func(*cobra.Command) error)
func (*cobra.Command).SetUsageTemplate(s string)
func (*cobra.Command).SetVersionTemplate(s string)
func (*cobra.Command).SuggestionsFor(typedName string) []string
func (*cobra.Command).Traverse(args []string) (*cobra.Command, []string, error)
func (*cobra.Command).Usage() error
func (*cobra.Command).UsageFunc() (f func(*cobra.Command) error)
func (*cobra.Command).UsagePadding() int
func (*cobra.Command).UsageString() string
func (*cobra.Command).UsageTemplate() string
func (*cobra.Command).UseLine() string
func (*cobra.Command).ValidateArgs(args []string) error
func (*cobra.Command).ValidateFlagGroups() error
func (*cobra.Command).ValidateRequiredFlags() error
func (*cobra.Command).VersionTemplate() string
func (*cobra.Command).VisitParents(fn func(*cobra.Command))
Command is just that, a command for your application. E.g. 'go run ...' - 'run' is the command. Cobra requires you to define the usage and description as part of your command definition to ensure usability.