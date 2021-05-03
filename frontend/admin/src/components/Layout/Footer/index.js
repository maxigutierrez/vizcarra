import React from 'react';
import { Layout } from 'antd';
import APPCONFIG from 'constants/appConfig';
const { Footer } = Layout;

const AppFooter = () => (
  <Footer className="app-footer app-footer-custom">
    <div className="footer-inner-v1">
      <span className="small">
        {/* eslint-disable-next-line */}
        © {APPCONFIG.year} <a className="brand" target="_blank">{APPCONFIG.brand}</a>
      </span>
     
    </div>
  </Footer>
)

export default AppFooter;
