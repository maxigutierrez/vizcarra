import APPCONFIG from 'constants/appConfig';

const colorOption = JSON.parse(localStorage.getItem("colorOption")) ? JSON.parse(localStorage.getItem("colorOption")) : APPCONFIG.settings.colorOption;
const initialSettings = {...APPCONFIG.settings, colorOption: colorOption};

const settings = (state = initialSettings, action) => {
  switch (action.type) {
    case "CHANGE_LAYOUT":
      return {
        ...state,
        layout: action.layoutOption
      };
    case "TOGGLE_BOXED_LAYOUT":
      return {
        ...state,
        boxedLayout: action.isBoxedLayout
      };
    case "TOGGLE_FIXED_SIDENAV":
      return {
        ...state,
        fixedSidenav: action.isFixedSidenav
      };
    case "TOGGLE_FIXED_HEADER":
      return {
        ...state,
        fixedHeader: action.isFixedHeader
      };
    case "TOGGLE_COLLAPSED_NAV":
      return {
        ...state,
        collapsedNav: action.isCollapsedNav
      };
    case "TOGGLE_OFFCANVAS_NAV":
      return {
        ...state,
        offCanvasNav: action.isOffCanvasNav
      };
    case "CHANGE_SIDENAV_WIDTH":
      return {
        ...state,
        sidenavWidth: action.sidenavWidth
      };
    case "TOGGLE_OFFCANVAS_MOBILE_NAV":
      return {
        ...state,
        offCanvasMobileNav: action.isOffCanvasMobileNav
      };
    case "CHANGE_COLOR_OPTION":
      localStorage.setItem("colorOption", JSON.stringify(action.colorOption));
      return {
        ...state,
        colorOption: action.colorOption
      };
    default:
      return state;
  }
}

export default settings;
