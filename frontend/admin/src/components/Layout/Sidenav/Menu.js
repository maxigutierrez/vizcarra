import React from 'react';
import { connect } from 'react-redux';
import { Menu, Icon } from 'antd';
import { toggleOffCanvasMobileNav } from 'actions/settings';
import _ from 'lodash'

class AppMenu extends React.Component {
  // eslint-disable-next-line
  constructor(props) {
    super(props)
  }

  // list for AccordionNav
  rootMenuItemKeys = [ // without submenu
    '/app/logistica',
  ]
  rootSubmenuKeys = [
    '/app/clientes',

  ];

  state = {
    openKeys: [],
  };

  onOpenChange = (openKeys) => {
    // AccordionNav
    // console.log(openKeys)
    const latestOpenKey = openKeys.find(key => this.state.openKeys.indexOf(key) === -1);
    if (this.rootSubmenuKeys.indexOf(latestOpenKey) === -1) {
      this.setState({ openKeys });
    } else {
      this.setState({
        openKeys: latestOpenKey ? [latestOpenKey] : [],
      });
    }
  }

  onMenuItemClick = (item) => {
    // AccordionNav
    const itemKey = item.key;
    if (this.rootMenuItemKeys.indexOf(itemKey) >= 0) {
      this.setState({ openKeys: [itemKey] });
    }

    const { isMobileNav } = this.props;
    if (isMobileNav) {
      this.closeMobileSidenav();
    }
  }

  closeMobileSidenav = () => {
    const { handleToggleOffCanvasMobileNav } = this.props;
    handleToggleOffCanvasMobileNav(true);
  }

  getSubMenuOrItem = item => {
    return <Menu.Item key={item.id}>
      <a href={"#/app/" + item.path}>
        <span>
          <Icon type={item.icon} /> <span className="nav-text">  {item.name}</span>
        </span>
      </a>
    </Menu.Item>;
  };

  getNavMenuItems = menusData => {
    if (!menusData) {
      return [];
    }
    return menusData
      .filter(item => !item.hideInMenu)
      .map(item => {
        // make dom
        const ItemDom = this.getSubMenuOrItem(item);
        return ItemDom;
      })
      .filter(item => item);
  }


  render() {
    const { collapsedNav, colorOption, location } = this.props;
    // const mode = collapsedNav ? 'vertical' : 'inline';
    const menuTheme = ['31', '32', '33', '34', '35', '36'].indexOf(colorOption) >= 0 ? 'light' : 'dark';
    const currentPathname = location.pathname;

    const menuProps = collapsedNav
      ? {}
      : {
        openKeys: this.state.openKeys
      };

      return (
        <Menu
          theme={menuTheme}
          mode="inline"
          inlineCollapsed={collapsedNav}
          {...menuProps}
          onOpenChange={this.onOpenChange}
          onClick={this.onMenuItemClick}
          // selectedKeys={[currentPathname]}
        >
          <Menu.Item key="/app/productos">
            <a href="#/app/productos">
              <span><Icon type="file"/><span className="nav-text">Productos</span></span>
            </a>
          </Menu.Item>
        </Menu>
      )
  }
}

const mapStateToProps = state => {
  // console.log(state);
  return ({
    collapsedNav: state.settings.collapsedNav,
    colorOption: state.settings.colorOption,
    location: state.routing.location,
    user: state.user,
  })
};

const mapDispatchToProps = dispatch => ({
  handleToggleOffCanvasMobileNav: (isOffCanvasMobileNav) => {
    dispatch(toggleOffCanvasMobileNav(isOffCanvasMobileNav));
  }
});

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(AppMenu);
