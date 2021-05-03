export function changeLayout(layoutOption) {
  return { type: "CHANGE_LAYOUT", layoutOption };
}
export function toggleBoxedLayout(isBoxedLayout) {
  return { type: "TOGGLE_BOXED_LAYOUT", isBoxedLayout: isBoxedLayout }
}
export function toggleCollapsedNav(isCollapsedNav) {
  return { type: "TOGGLE_COLLAPSED_NAV", isCollapsedNav: isCollapsedNav }
}
export function toggleOffCanvasNav(isOffCanvasNav) {
  return { type: "TOGGLE_OFFCANVAS_NAV", isOffCanvasNav: isOffCanvasNav }
}
export function toggleFixedSidenav(isFixedSidenav) {
  return { type: "TOGGLE_FIXED_SIDENAV", isFixedSidenav: isFixedSidenav }
}
export function toggleFixedHeader(isFixedHeader) {
  return { type: "TOGGLE_FIXED_HEADER", isFixedHeader: isFixedHeader }
}
export function changeSidenavWidth(sidenavWidth) {
  return { type: "CHANGE_SIDENAV_WIDTH", sidenavWidth: sidenavWidth }
}
export function toggleOffCanvasMobileNav(isOffCanvasMobileNav) {
  return { type: "TOGGLE_OFFCANVAS_MOBILE_NAV", isOffCanvasMobileNav: isOffCanvasMobileNav }
}
export function changeColorOption(colorOption) {
  return { type: "CHANGE_COLOR_OPTION", colorOption: colorOption }
}
export function changeTheme(themeOption) {
  return { type: "CHANGE_THEME", theme: themeOption }
}
