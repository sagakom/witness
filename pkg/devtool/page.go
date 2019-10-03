package devtool

// NavigationResult https://chromedevtools.github.io/devtools-protocol/tot/Page#method-navigate
type NavigationResult struct {
	FrameID   string `json:"frameId"`
	LoaderID  string `json:"loaderId"`
	ErrorText string `json:"errorText"`
}

// NavigationEntry https://chromedevtools.github.io/devtools-protocol/tot/Page#type-NavigationEntry
type NavigationEntry struct {
	ID             int64  `json:"id"`
	URL            string `json:"url"`
	UserTypedURL   string `json:"userTypedURL"`
	Title          string `json:"title"`
	TransitionType string `json:"transitionType"`
}

// LayoutViewport https://chromedevtools.github.io/devtools-protocol/tot/Page#type-LayoutViewport
type LayoutViewport struct {
	PageX        int64 `json:"pageX"`
	PageY        int64 `json:"pageY"`
	ClientWidth  int64 `json:"clientWidth"`
	ClientHeight int64 `json:"clientHeight"`
}

// Viewport https://chromedevtools.github.io/devtools-protocol/tot/Page#type-Viewport
type Viewport struct {
	X      float64 `json:"x"`
	Y      float64 `json:"y"`
	Width  float64 `json:"width"`
	Height float64 `json:"height"`
	Scale  float64 `json:"scale"`
}

// VisualViewport https://chromedevtools.github.io/devtools-protocol/tot/Page#type-VisualViewport
type VisualViewport struct {
	OffsetX      float64 `json:"offsetX"`
	OffsetY      float64 `json:"offsetY"`
	PageX        float64 `json:"pageX"`
	PageY        float64 `json:"pageY"`
	ClientWidth  float64 `json:"clientWidth"`
	ClientHeight float64 `json:"clientHeight"`
	Scale        float64 `json:"scale"`
	Zoom         float64 `json:"zoom"`
}

// LayoutMetrics https://chromedevtools.github.io/devtools-protocol/tot/Page#method-getLayoutMetrics
type LayoutMetrics struct {
	LayoutViewport *LayoutViewport `json:"layoutViewport"`
	VisualViewport *VisualViewport `json:"visualViewport"`
	ContentSize    *Rect           `json:"contentSize"`
}

// NavigationHistory https://chromedevtools.github.io/devtools-protocol/tot/Page#method-getNavigationHistory
type NavigationHistory struct {
	CurrentIndex int64              `json:"currentIndex"`
	Entries      []*NavigationEntry `json:"entries"`
}