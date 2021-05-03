import React from 'react';
import { connect } from 'react-redux';
// import APPCONFIG from 'constants/appConfig';
import App from './App'
// import ContentOnly from './ContentOnly'
// import HeaderContentFooter from './HeaderContentFooter'

class AppLayout extends React.Component {

  updateLayout(layout, boxedLayout, fixedSidenav, fixedHeader) {
      return <App boxedLayout={boxedLayout} fixedSidenav={fixedSidenav} fixedHeader={fixedHeader} />;
  }

  render() {
    const { layout, boxedLayout, fixedSidenav, fixedHeader } = this.props;

    return (
      <div id="app-layout-container">
        { this.updateLayout(layout, boxedLayout, fixedSidenav, fixedHeader) }
      </div>
    )
  }
}

const mapStateToProps = (state, ownProps) => ({
  layout: state.settings.layout,
  boxedLayout: state.settings.boxedLayout,
  fixedSidenav: state.settings.fixedSidenav,
  fixedHeader: state.settings.fixedHeader
});

export default connect(
  mapStateToProps
)(AppLayout);
