import React from 'react';
import { Route, Redirect } from 'react-router-dom';
import loadable from 'react-loadable';
import LoadingComponent from 'components/Loading';

// 3rd
import 'styles/antd.less';
import 'styles/bootstrap/bootstrap.scss'
// custom
import "styles/layout.scss"
import "styles/theme.scss"
import "styles/ui.scss"
import "styles/vendors.scss"

let AsyncAppLayout = loadable({
  loader: () => import('components/Layout/AppLayout/'),
  loading: LoadingComponent
})

// let AsyncAccount = loadable({
//   loader: () => import('routes/login/'),
//   loading: LoadingComponent
// })

class App extends React.Component {
  render() {
    const { match, location } = this.props;
    const isRoot = location.pathname === '/' ? true : false;

    if (isRoot) {
      return ( <Redirect to={'/app'}/> );
    }

    return (
      <div id="app">
        {/* <Route path={`${match.url}login`} component={AsyncAccount} /> */}
        <Route path={`${match.url}app`} component={AsyncAppLayout} />
      </div>
    );
  }
}

export default App;
