import React from 'react';
import './styles.scss';
import { api } from 'api';
import { connect } from 'react-redux';
import { withRouter } from "react-router-dom";
import DEMO from 'constants/demoData';
import { message, Layout, Menu, Dropdown, Avatar, Divider, Icon } from 'antd';
import Logo from 'components/Logo';
import { toggleCollapsedNav, toggleOffCanvasMobileNav } from 'actions/settings';
import { setUser } from 'actions/user';

const { Header } = Layout;

class AppHeader extends React.Component {

  constructor(props) {
    super(props)
    this.state = {
      openChangePassword: false,
      disabledLogin: false,
      confirmLoading: false,
    }
  }

  async componentDidMount() {
    // try {
    //   const response = await api.auth.getAuthenticatedUser();
    //   if (response.status === "success") {
    //     this.props.handleSetUser(response.data.usuario);
    //     this.setState({ openChangePassword: response.data.usuario.cambiarcontraseña });
    //     this.fetch();
    //   } else {
    //     this.props.history.push('/login');
    //   }
    // } catch (e) {
    //   message.error(e.toString(), 10);
    // }
  }

  fetch = async (params = {}) => {
  }

  onToggleCollapsedNav = () => {
    const { handleToggleCollapsedNav, collapsedNav } = this.props;
    handleToggleCollapsedNav(!collapsedNav);
  }

  onToggleOffCanvasMobileNav = () => {
    const { handleToggleOffCanvasMobileNav, offCanvasMobileNav } = this.props;
    handleToggleOffCanvasMobileNav(!offCanvasMobileNav);
  }

  render() {
    const { collapsedNav, showLogo } = this.props;

    return (
      <Header className="app-header">
        <div
          className='app-header-inner' style={{ backgroundColor: '#001529' }}
        >
          <div className="header-left">
            <div className="list-unstyled list-inline">
              {showLogo && [
                <Logo key="logo" />,
                <Divider type="vertical" key="line" />,
              ]}
              <a href={DEMO.link} className="list-inline-item d-none d-md-inline-block" onClick={this.onToggleCollapsedNav}>
                <Icon type={collapsedNav ? 'menu-unfold' : 'menu-fold'} className="list-icon" style={{ color: this.state.mouseOver ? '#1a25ab' : '#d4d4d4' }} onMouseOver={() => this.setState({ mouseOver: true })} onMouseOut={() => this.setState({ mouseOver: false })} />
              </a>
            </div>
          </div>
          {/* <div className="header-right">
            <div className="list-unstyled list-inline">
              <Dropdown className="list-inline-item" overlay={
                <Menu className="app-header-dropdown">
                  <Menu.Item key="4" className="d-block d-md-none"> Signed in as <strong>{this.props.user && this.props.user.nombre}</strong> </Menu.Item>
                  <Menu.Divider className="d-block d-md-none" />
                  <Menu.Item key="3" onClick={() => {
                    this.props.history.push('/login');
                  }}> <a href="/"><Icon type="logout" />Salir</a> </Menu.Item>
                </Menu>
              } trigger={['click']} placement="bottomRight">
                <a className="ant-dropdown-link no-link-style" href={DEMO.link}>
                  <Avatar size="small">{this.props.user && this.props.user.nombre[0]}</Avatar>
                  <span style={{ color: '#d4d4d4' }} className="avatar-text d-none d-md-inline">{this.props.user && this.props.user.nombre}</span>
                </a>
              </Dropdown>
            </div>
          </div> */}
        </div>
      </Header>
    );
  }
}

const mapStateToProps = (state) => ({
  offCanvasMobileNav: state.settings.offCanvasMobileNav,
  collapsedNav: state.settings.collapsedNav,
  colorOption: state.settings.colorOption,
  user: state.user,
});

const mapDispatchToProps = dispatch => ({
  handleToggleCollapsedNav: (isCollapsedNav) => {
    dispatch(toggleCollapsedNav(isCollapsedNav));
  },
  handleToggleOffCanvasMobileNav: (isOffCanvasMobileNav) => {
    dispatch(toggleOffCanvasMobileNav(isOffCanvasMobileNav));
  },
  handleSetUser: (user) => {
    dispatch(setUser(user));
  }
});


const WrappedAppHeader = withRouter(AppHeader);

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(WrappedAppHeader);
