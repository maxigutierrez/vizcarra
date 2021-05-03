import React from 'react';
import { Route } from 'react-router-dom';
import { withRouter } from 'react-router'
import loadable from 'react-loadable';
import LoadingComponent from 'components/Loading';
import { Layout } from 'antd';
const { Content } = Layout;

let AsyncDashboard = loadable({
  loader: () => import('routes/dashboard/'),
  loading: LoadingComponent
})
let AsyncProductos = loadable({
  loader: () => import('routes/productos'),
  loading: LoadingComponent,
})


class AppContent extends React.Component {
  render() {
    const { match } = this.props;

    return (
      <Content id='app-content'>
        <Route exact path={`${match.url}/dashboard`} component={AsyncDashboard} />
        <Route exact path={`${match.url}/productos`} component={AsyncProductos} />
      </Content>
    );
  }
}
export default withRouter(AppContent);