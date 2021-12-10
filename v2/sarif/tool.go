package sarif

// Tool ...
type Tool struct {
	PropertyBag
	Driver *ToolComponent `json:"driver"`
}

// NewTool creates a new Tool and returns a pointer to it
func NewTool(driver *ToolComponent) *Tool {
	return &Tool{
		Driver: driver,
	}
}

// NewSimpleTool creates a new SimpleTool and returns a pointer to it
func NewSimpleTool(driverName string) *Tool {
	return &Tool{
		Driver: NewDriver(driverName),
	}
}
