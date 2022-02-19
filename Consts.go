package implot

import "C"
import "github.com/inkyblackness/imgui-go/v4"

// Special auto value. Used in sizes, width, etc.
var Auto float32 = -1

// Special color used to indicate that a color should be deduced automatically.
var AutoColor = imgui.Vec4{X: 0, Y: 0, Z: 0, W: -1}

type Axis C.int           // Axis indices
type Flags C.int          // Flags for plots / BeginPlot
type AxisFlags C.int      // Flags for plot axes / SetupAxis
type SubplotFlags C.int   // Flags for subplots / BeginSubplot
type LegendFlags C.int    // Flags for legends / SetupLegend
type MouseTextFlags C.int // Flags for mouse hover text / SetupMouseText
type DragToolFlags C.int  // Flags for DragPoint, DragLine, DragRect
type BarGroupsFlags C.int // Flags for PlotBarGroups

type Condition C.int // Represents a condition for SetupAxisLimits etc. (a subset of imgui.Cond)
type StyleCol C.int  // Plot styling colors
type StyleVar C.int  // Plot styling variables
type Marker C.int    // Markers
type Colormap C.int  // Built-in colormaps
type Location C.int  // Locations used to position items on a plot
type Bin C.int       // Different automatic histogram binning methods

// Axis indices. The values assigned may change; NEVER hardcode these.
const (
	Axis_X1 Axis = iota // enabled by default
	Axis_X2             // diabled by default
	Axis_X3             // diabled by default
	Axis_Y1             // enabled by default
	Axis_Y2             // diabled by default
	Axis_Y3             // diabled by default
	Axis_Count
)

// Flags for plots / BeginPlot
const (
	Flags_NoTitle     Flags = 1 << iota // the plot title will not be displayed (titles are also hidden if preceeded by double hashes, e.g. "##MyPlot")
	Flags_NoLegend                      // the legend will not be displayed
	Flags_NoMouseText                   // the mouse position, in plot coordinates, will not be displayed inside of the plot
	Flags_NoInputs                      // the user will not be able to interact with the plot
	Flags_NoMenus                       // the user will not be able to open context menus
	Flags_NoBoxSelect                   // the user will not be able to box-select
	Flags_NoChild                       // a child window region will not be used to capture mouse scroll (can boost performance for single ImGui window applications)
	Flags_NoFrame                       // the ImGui frame will not be rendered
	Flags_Equal                         // x and y axes pairs will be constrained to have the same units/pixel
	Flags_Crosshairs                    // the default mouse cursor will be replaced with a crosshair when hovered
	Flags_AntiAliased                   // plot items will be software anti-aliased (not recommended for high density plots, prefer MSAA)

	Flags_None       = 0 // default
	Flags_CanvasOnly = Flags_NoTitle | Flags_NoLegend | Flags_NoMenus | Flags_NoBoxSelect | Flags_NoMouseText
)

// Flags for plot axes / SetupAxis
const (
	AxisFlags_NoLabel      AxisFlags = 1 << iota // the axis label will not be displayed (axis labels also hidden if the supplied string name is NULL)
	AxisFlags_NoGridLines                        // no grid lines will be displayed
	AxisFlags_NoTickMarks                        // no tick marks will be displayed
	AxisFlags_NoTickLabels                       // no text labels will be displayed
	AxisFlags_NoInitialFit                       // axis will not be initially fit to data extents on the first rendered frame
	AxisFlags_NoMenus                            // the user will not be able to open context menus with right-click
	AxisFlags_Opposite                           // axis ticks and labels will be rendered on conventionally opposite side (i.e, right or top)
	AxisFlags_Foreground                         // grid lines will be displayed in the foreground (i.e. on top of data) in stead of the background
	AxisFlags_LogScale                           // a logartithmic (base 10) axis scale will be used (mutually exclusive with ImPlotAxisFlags_Time)
	AxisFlags_Time                               // axis will display date/time formatted labels (mutually exclusive with ImPlotAxisFlags_LogScale)
	AxisFlags_Invert                             // the axis will be inverted
	AxisFlags_AutoFit                            // axis will be auto-fitting to data extents
	AxisFlags_RangeFit                           // axis will only fit points if the point is in the visible range of the **orthogonal** axis
	AxisFlags_LockMin                            // the axis minimum value will be locked when panning/zooming
	AxisFlags_LockMax                            // the axis maximum value will be locked when panning/zooming

	AxisFlags_None          = 0 // default
	AxisFlags_Lock          = AxisFlags_LockMin | AxisFlags_LockMax
	AxisFlags_NoDecorations = AxisFlags_NoLabel | AxisFlags_NoGridLines | AxisFlags_NoTickMarks | AxisFlags_NoTickLabels
	AxisFlags_AuxDefault    = AxisFlags_NoGridLines | AxisFlags_Opposite
)

// Flags for subplots / BeginSubplot
const (
	SubplotFlags_NoTitle    SubplotFlags = 1 << iota // the subplot title will not be displayed (titles are also hidden if preceeded by double hashes, e.g. "##MySubplot")
	SubplotFlags_NoLegend                            // the legend will not be displayed (only applicable if ImPlotSubplotFlags_ShareItems is enabled)
	SubplotFlags_NoMenus                             // the user will not be able to open context menus with right-click
	SubplotFlags_NoResize                            // resize splitters between subplot cells will be not be provided
	SubplotFlags_NoAlign                             // subplot edges will not be aligned vertically or horizontally
	SubplotFlags_ShareItems                          // items across all subplots will be shared and rendered into a single legend entry
	SubplotFlags_LinkRows                            // link the y-axis limits of all plots in each row (does not apply to auxiliary axes)
	SubplotFlags_LinkCols                            // link the x-axis limits of all plots in each column (does not apply to auxiliary axes)
	SubplotFlags_LinkAllX                            // link the x-axis limits in every plot in the subplot (does not apply to auxiliary axes)
	SubplotFlags_LinkAllY                            // link the y-axis limits in every plot in the subplot (does not apply to auxiliary axes)
	SubplotFlags_ColMajor                            // subplots are added in column major order instead of the default row major order
	SubplotFlags_None       = 0                      // default
)

// Flags for legends / SetupLegend
const (
	LegendFlags_NoButtons       LegendFlags = 1 << iota // legend icons will not function as hide/show buttons
	LegendFlags_NoHighlightItem                         // plot items will not be highlighted when their legend entry is hovered
	LegendFlags_NoHighlightAxis                         // axes will not be highlighted when legend entries are hovered (only relevant if x/y-axis count > 1)
	LegendFlags_NoMenus                                 // the user will not be able to open context menus with right-click
	LegendFlags_Outside                                 // legend will be rendered outside of the plot area
	LegendFlags_Horizontal                              // legend entries will be displayed horizontally
	LegendFlags_None            = 0                     // default
)

// Flags for mouse hover text / SetupMouseText
const (
	MouseTextFlags_NoAuxAxes  MouseTextFlags = 1 << iota // only show the mouse position for primary axes
	MouseTextFlags_NoFormat                              // axes label formatters won't be used to render text
	MouseTextFlags_ShowAlways                            // always display mouse position even if plot not hovered
	MouseTextFlags_None       = 0                        // default
)

// Flags for DragPoint, DragLine, DragRect
const (
	DragToolFlags_NoCursors DragToolFlags = 1 << iota // drag tools won't change cursor icons when hovered or held
	DragToolFlags_NoFit                               // the drag tool won't be considered for plot fits
	DragToolFlags_NoInputs                            // lock the tool from user inputs
	DragToolFlags_Delayed                             // tool rendering will be delayed one frame; useful when applying position-constraints
	DragToolFlags_None      = 0                       // default
)

// Flags for PlotBarGroups
const (
	BarGroupsFlags_Stacked BarGroupsFlags = 1 << iota // items in a group will be stacked on top of each other
	BarGroupsFlags_None                   = 0         // default
)

// Represents a condition for SetupAxisLimits etc. (a subset of imgui.Cond)
const (
	Condition_None   = Condition(imgui.ConditionNone)   // No condition (always set the variable), same as _Always
	Condition_Always = Condition(imgui.ConditionAlways) // No condition (always set the variable)
	Condition_Once   = Condition(imgui.ConditionOnce)   // Set the variable once per runtime session (only the first call will succeed)
)

// Plot styling colors
const (
	// item styling colors
	StyleCol_Line          StyleCol = iota // plot line/outline color (defaults to next unused color in current colormap)
	StyleCol_Fill                          // plot fill color for bars (defaults to the current line color)
	StyleCol_MarkerOutline                 // marker outline color (defaults to the current line color)
	StyleCol_MarkerFill                    // marker fill color (defaults to the current line color)
	StyleCol_ErrorBar                      // error bar color (defaults to ImGuiCol_Text)
	// plot styling colors
	StyleCol_FrameBg       // plot frame background color (defaults to ImGuiCol_FrameBg)
	StyleCol_PlotBg        // plot area background color (defaults to ImGuiCol_WindowBg)
	StyleCol_PlotBorder    // plot area border color (defaults to ImGuiCol_Border)
	StyleCol_LegendBg      // legend background color (defaults to ImGuiCol_PopupBg)
	StyleCol_LegendBorder  // legend border color (defaults to ImPlotCol_PlotBorder)
	StyleCol_LegendText    // legend text color (defaults to ImPlotCol_InlayText)
	StyleCol_TitleText     // plot title text color (defaults to ImGuiCol_Text)
	StyleCol_InlayText     // color of text appearing inside of plots (defaults to ImGuiCol_Text)
	StyleCol_AxisText      // axis label and tick lables color (defaults to ImGuiCol_Text)
	StyleCol_AxisGrid      // axis grid color (defaults to 25% ImPlotCol_AxisText)
	StyleCol_AxisTick      // axis tick color (defaults to AxisGrid)
	StyleCol_AxisBg        // background color of axis hover region (defaults to transparent)
	StyleCol_AxisBgHovered // axis hover color (defaults to ImGuiCol_ButtonHovered)
	StyleCol_AxisBgActive  // axis active color (defaults to ImGuiCol_ButtonActive)
	StyleCol_Selection     // box-selection color (defaults to yellow)
	StyleCol_Crosshairs    // crosshairs color (defaults to ImPlotCol_PlotBorder)
	StyleCol_Count
)

// Plot styling variables
const (
	// item styling variables
	StyleVar_LineWeight       StyleVar = iota // float,  plot item line weight in pixels
	StyleVar_Marker                           // int,    marker specification
	StyleVar_MarkerSize                       // float,  marker size in pixels (roughly the marker's "radius")
	StyleVar_MarkerWeight                     // float,  plot outline weight of markers in pixels
	StyleVar_FillAlpha                        // float,  alpha modifier applied to all plot item fills
	StyleVar_ErrorBarSize                     // float,  error bar whisker width in pixels
	StyleVar_ErrorBarWeight                   // float,  error bar whisker weight in pixels
	StyleVar_DigitalBitHeight                 // float,  digital channels bit height (at 1) in pixels
	StyleVar_DigitalBitGap                    // float,  digital channels bit padding gap in pixels
	// plot styling variables
	StyleVar_PlotBorderSize     // float,  thickness of border around plot area
	StyleVar_MinorAlpha         // float,  alpha multiplier applied to minor axis grid lines
	StyleVar_MajorTickLen       // ImVec2, major tick lengths for X and Y axes
	StyleVar_MinorTickLen       // ImVec2, minor tick lengths for X and Y axes
	StyleVar_MajorTickSize      // ImVec2, line thickness of major ticks
	StyleVar_MinorTickSize      // ImVec2, line thickness of minor ticks
	StyleVar_MajorGridSize      // ImVec2, line thickness of major grid lines
	StyleVar_MinorGridSize      // ImVec2, line thickness of minor grid lines
	StyleVar_PlotPadding        // ImVec2, padding between widget frame and plot area, labels, or outside legends (i.e. main padding)
	StyleVar_LabelPadding       // ImVec2, padding between axes labels, tick labels, and plot edge
	StyleVar_LegendPadding      // ImVec2, legend padding from plot edges
	StyleVar_LegendInnerPadding // ImVec2, legend inner padding from legend edges
	StyleVar_LegendSpacing      // ImVec2, spacing between legend entries
	StyleVar_MousePosPadding    // ImVec2, padding between plot edge and interior info text
	StyleVar_AnnotationPadding  // ImVec2, text padding around annotation labels
	StyleVar_FitPadding         // ImVec2, additional fit padding as a percentage of the fit extents (e.g. ImVec2(0.1f,0.1f) adds 10% to the fit extents of X and Y)
	StyleVar_PlotDefaultSize    // ImVec2, default size used when ImVec2(0,0) is passed to BeginPlot
	StyleVar_PlotMinSize        // ImVec2, minimum size plot frame can be when shrunk
	StyleVar_Count
)

// Markers
const (
	Marker_Circle   Marker = iota // a circle marker
	Marker_Square                 // a square maker
	Marker_Diamond                // a diamond marker
	Marker_Up                     // an upward-pointing triangle marker
	Marker_Down                   // an downward-pointing triangle marker
	Marker_Left                   // an leftward-pointing triangle marker
	Marker_Right                  // an rightward-pointing triangle marker
	Marker_Cross                  // a cross marker (not fillable)
	Marker_Plus                   // a plus marker (not fillable)
	Marker_Asterisk               // a asterisk marker (not fillable)
	Marker_Count
	Marker_None = -1 // no marker
)

// Built-in colormaps
const (
	Colormap_Deep     Colormap = iota // a.k.a. seaborn deep             (qual=true,  n=10) (default)
	Colormap_Dark                     // a.k.a. matplotlib "Set1"        (qual=true,  n=9 )
	Colormap_Pastel                   // a.k.a. matplotlib "Pastel1"     (qual=true,  n=9 )
	Colormap_Paired                   // a.k.a. matplotlib "Paired"      (qual=true,  n=12)
	Colormap_Viridis                  // a.k.a. matplotlib "viridis"     (qual=false, n=11)
	Colormap_Plasma                   // a.k.a. matplotlib "plasma"      (qual=false, n=11)
	Colormap_Hot                      // a.k.a. matplotlib/MATLAB "hot"  (qual=false, n=11)
	Colormap_Cool                     // a.k.a. matplotlib/MATLAB "cool" (qual=false, n=11)
	Colormap_Pink                     // a.k.a. matplotlib/MATLAB "pink" (qual=false, n=11)
	Colormap_Jet                      // a.k.a. MATLAB "jet"             (qual=false, n=11)
	Colormap_Twilight                 // a.k.a. matplotlib "twilight"    (qual=false, n=11)
	Colormap_RdBu                     // red/blue, Color Brewer          (qual=false, n=11)
	Colormap_BrBG                     // brown/blue-green, Color Brewer  (qual=false, n=11)
	Colormap_PiYG                     // pink/yellow-green, Color Brewer (qual=false, n=11)
	Colormap_Spectral                 // color spectrum, Color Brewer    (qual=false, n=11)
	Colormap_Greys                    // white/black                     (qual=false, n=2 )
)

// Locations used to position items on a plot
const (
	Location_North  Location = 1 << iota // top-center
	Location_South                       // bottom-center
	Location_West                        // center-left
	Location_East                        // center-right
	Location_Center = 0                  // center-center

	Location_NorthWest = Location_North | Location_West // top-left
	Location_NorthEast = Location_North | Location_East // top-right
	Location_SouthWest = Location_South | Location_West // bottom-left
	Location_SouthEast = Location_South | Location_East // bottom-right
)

// Different automatic histogram binning methods (k = bin count or w = bin width)
const (
	Bin_Sqrt    Bin = -1 // k = sqrt(n)
	Bin_Sturges     = -2 // k = 1 + log2(n)
	Bin_Rice        = -3 // k = 2 * cbrt(n)
	Bin_Scott       = -4 // w = 3.49 * sigma / cbrt(n)
)
