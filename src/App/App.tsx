import React from 'react';
import SideNavigation from './SideNavigation/SideNavigation';
import Header from './Header/Header';
import Subheader from './SubHeader/Subheader';
import Footer from './Footer/Footer';
import HeaderMobile from './HeaderMobile/HeaderMobile';

function App() {
  return (
    <>
    <HeaderMobile/>
      <div className="kt-grid kt-grid--hor kt-grid--root">
        <div className="kt-grid__item kt-grid__item--fluid kt-grid kt-grid--ver kt-page">
        <SideNavigation />

          <div className="kt-grid__item kt-grid__item--fluid kt-grid kt-grid--hor kt-wrapper" id="kt_wrapper">
            <Header/>
            <div className="kt-content  kt-grid__item kt-grid__item--fluid kt-grid kt-grid--hor" id="kt_content">
              <Subheader/>
              <div className="kt-container  kt-container--fluid  kt-grid__item kt-grid__item--fluid">
                <div className="row"></div>
              </div>
              
            </div>
          </div>
        </div>
      </div>
    </>
  );
}

export default App;
