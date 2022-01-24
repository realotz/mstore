const setting = {
  showSettingButton: true,
  showDarkModeToggle: true,
  settingButtonPosition: 'auto',
  permissionMode: 'ROUTE_MAPPING',
  permissionCacheType: 1,
  sessionTimeoutProcessing: 0,
  themeColor: '#0084f4',
  grayMode: false,
  colorWeak: false,
  fullContent: false,
  contentMode: 'full',
  showLogo: true,
  showFooter: false,

  // Header configuration
  headerSetting: {
    bgColor: '#394664',
    fixed: true,
    show: true,
    theme: 'dark',
    useLockPage: true,
    showFullScreen: true,
    showDoc: true,
    showNotice: true,
    showSearch: false,
  },

  // Menu configuration
  menuSetting: {
    bgColor: '#ffffff',
    fixed: true,
    collapsed: false,
    collapsedShowTitle: false,
    canDrag: false,
    show: true,
    hidden: false,
    menuWidth: 210,
    mode: 'inline',
    type: 'mix',
    theme: 'light',
    split: false,
    topMenuAlign: 'center',
    trigger: 'FOOTER',
    accordion: true,
    closeMixSidebarOnChange: false,
    mixSideTrigger: 'click',
    mixSideFixed: false,
  },

  // Multi-label
  multiTabsSetting: {
    cache: false,
    show: true,
    canDrag: true,
    showQuick: true,
    showRedo: true,
    showFold: true,
  },

  // Transition Setting
  transitionSetting: {
    enable: true,
    basicTransition: 'fade-slide',
    openPageLoading: true,
    openNProgress: false,
  },
  openKeepAlive: true,
  lockTime: 0,
  showBreadCrumb: true,
  showBreadCrumbIcon: true,
  useErrorHandle: false,
  useOpenBackTop: true,
  canEmbedIFramePage: true,
  closeMessageOnSwitch: true,
  removeAllHttpPending: false,
};

export default setting;
