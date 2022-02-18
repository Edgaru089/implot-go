 
## ImPlot for Go
This library is a [Go](https://www.golang.org) wrapper for [ImPlot](https://github.com/epezent/implot/), for use with [imgui-go](https://github.com/inkyblackness/imgui-go) which wraps [Dear ImGui](https://github.com/ocornut/imgui).

It currently targets `ImPlot v0.13` with `imgui-go v4.4.0` (which wraps `Dear ImGui v1.85`).

It has similar goals compared to inkyblackness's imgui-go wrapper:
 - [x] hand-crafted
 - [x] documented
 - [ ] feature-complete
 - [ ] versioned
 - [x] with examples (living in the `example-intree` branch for now)

![Screenshot](example/screenshot.png)

---

This is very much work in progress, here is a list from the Table of Contents of implot.h:
 - [x] Contexts
 - [x] Begin/End Plot
 - [x] Begin/End Subplot
 - [x] Setup
 - [ ] SetNext
 - [ ] Plot Items
 - [ ] Plot Tools
 - [ ] Plot Utils
 - [ ] Legend Utils
 - [ ] Drag and Drop
 - [ ] Styling
 - [ ] Colormaps
 - [ ] Input Mapping
 - [ ] Miscellaneous