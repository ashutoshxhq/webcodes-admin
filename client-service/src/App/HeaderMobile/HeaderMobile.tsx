import React from 'react'
import './HeaderMobile.css'

function HeaderMobile() {
    return (
        <div id="kt_header_mobile" className="kt-header-mobile  kt-header-mobile--fixed ">
        <div className="kt-header-mobile__logo">
            <a href="index.html">
                <img alt="Logo" src="/assets/media/logos/logo.png" className="h-30" />
            </a>
        </div>

        
        <div className="kt-header-mobile__toolbar">
            <button className="kt-header-mobile__toggler kt-header-mobile__toggler--left" id="kt_aside_mobile_toggler"><span></span></button>

            <div className="kt-header__topbar-item kt-header__topbar-item--search dropdown ml-4" id="kt_quick_search_toggle">
                            <div className="kt-header__topbar-wrapper" data-toggle="dropdown" data-offset="10px,0px">
                                <span className="kt-header__topbar-icon">
									<svg xmlns="http://www.w3.org/2000/svg" xmlnsXlink="http://www.w3.org/1999/xlink" width="24px" height="24px" viewBox="0 0 24 24" version="1.1" className="kt-svg-icon">
    <g stroke="none" stroke-width="1" fill="none" fill-rule="evenodd">
        <rect x="0" y="0" width="24" height="24"/>
        <path d="M14.2928932,16.7071068 C13.9023689,16.3165825 13.9023689,15.6834175 14.2928932,15.2928932 C14.6834175,14.9023689 15.3165825,14.9023689 15.7071068,15.2928932 L19.7071068,19.2928932 C20.0976311,19.6834175 20.0976311,20.3165825 19.7071068,20.7071068 C19.3165825,21.0976311 18.6834175,21.0976311 18.2928932,20.7071068 L14.2928932,16.7071068 Z" fill="#000000" fill-rule="nonzero" opacity="0.3"/>
        <path d="M11,16 C13.7614237,16 16,13.7614237 16,11 C16,8.23857625 13.7614237,6 11,6 C8.23857625,6 6,8.23857625 6,11 C6,13.7614237 8.23857625,16 11,16 Z M11,18 C7.13400675,18 4,14.8659932 4,11 C4,7.13400675 7.13400675,4 11,4 C14.8659932,4 18,7.13400675 18,11 C18,14.8659932 14.8659932,18 11,18 Z" fill="#000000" fill-rule="nonzero"/>
    </g>
</svg>							</span>
                            </div>
                            <div className="dropdown-menu dropdown-menu-fit dropdown-menu-right dropdown-menu-anim dropdown-menu-lg">
                                <div className="kt-quick-search kt-quick-search--dropdown kt-quick-search--result-compact" id="kt_quick_search_dropdown">
                                    <form method="get" className="kt-quick-search__form">
                                        <div className="input-group">
                                            <div className="input-group-prepend"><span className="input-group-text"><i className="flaticon2-search-1"></i></span></div>
                                            <input type="text" className="form-control kt-quick-search__input" placeholder="Search..."/>
                                            <div className="input-group-append"><span className="input-group-text"><i className="la la-close kt-quick-search__close"></i></span></div>
                                        </div>
                                    </form>
                                    <div className="kt-quick-search__wrapper kt-scroll" data-scroll="true" data-height="325" data-mobile-height="200">

                                    </div>
                                </div>
                            </div>
                        </div>
            <div className="kt-header__topbar-item dropdown ml-4">
                            <div className="kt-header__topbar-wrapper" data-toggle="dropdown" data-offset="30px,0px" aria-expanded="true">
                                <span className="kt-header__topbar-icon kt-pulse kt-pulse--brand">
                                    <svg xmlns="http://www.w3.org/2000/svg" xmlnsXlink="http://www.w3.org/1999/xlink" width="24px" height="24px" viewBox="0 0 24 24" version="1.1" className="kt-svg-icon">
    <g stroke="none" stroke-width="1" fill="none" fill-rule="evenodd">
        <rect x="0" y="0" width="24" height="24"/>
        <path d="M2.56066017,10.6819805 L4.68198052,8.56066017 C5.26776695,7.97487373 6.21751442,7.97487373 6.80330086,8.56066017 L8.9246212,10.6819805 C9.51040764,11.267767 9.51040764,12.2175144 8.9246212,12.8033009 L6.80330086,14.9246212 C6.21751442,15.5104076 5.26776695,15.5104076 4.68198052,14.9246212 L2.56066017,12.8033009 C1.97487373,12.2175144 1.97487373,11.267767 2.56066017,10.6819805 Z M14.5606602,10.6819805 L16.6819805,8.56066017 C17.267767,7.97487373 18.2175144,7.97487373 18.8033009,8.56066017 L20.9246212,10.6819805 C21.5104076,11.267767 21.5104076,12.2175144 20.9246212,12.8033009 L18.8033009,14.9246212 C18.2175144,15.5104076 17.267767,15.5104076 16.6819805,14.9246212 L14.5606602,12.8033009 C13.9748737,12.2175144 13.9748737,11.267767 14.5606602,10.6819805 Z" fill="#000000" opacity="0.3"/>
        <path d="M8.56066017,16.6819805 L10.6819805,14.5606602 C11.267767,13.9748737 12.2175144,13.9748737 12.8033009,14.5606602 L14.9246212,16.6819805 C15.5104076,17.267767 15.5104076,18.2175144 14.9246212,18.8033009 L12.8033009,20.9246212 C12.2175144,21.5104076 11.267767,21.5104076 10.6819805,20.9246212 L8.56066017,18.8033009 C7.97487373,18.2175144 7.97487373,17.267767 8.56066017,16.6819805 Z M8.56066017,4.68198052 L10.6819805,2.56066017 C11.267767,1.97487373 12.2175144,1.97487373 12.8033009,2.56066017 L14.9246212,4.68198052 C15.5104076,5.26776695 15.5104076,6.21751442 14.9246212,6.80330086 L12.8033009,8.9246212 C12.2175144,9.51040764 11.267767,9.51040764 10.6819805,8.9246212 L8.56066017,6.80330086 C7.97487373,6.21751442 7.97487373,5.26776695 8.56066017,4.68198052 Z" fill="#000000"/>
    </g>
</svg>                                
                                </span>
                            </div>
                            <div className="dropdown-menu dropdown-menu-fit dropdown-menu-right dropdown-menu-anim dropdown-menu-top-unround dropdown-menu-lg">
                                <form>
                                  
                                    <div className="tab-content">
                                        <div className="tab-pane active show" id="topbar_notifications_notifications" role="tabpanel">
                                            <div className="kt-notification kt-margin-t-10 kt-margin-b-10 kt-scroll" data-scroll="true" data-height="300" data-mobile-height="200">
                                                <a href="#" className="kt-notification__item">
                                                    <div className="kt-notification__item-icon">
                                                        <i className="flaticon2-line-chart kt-font-success"></i>
                                                    </div>
                                                    <div className="kt-notification__item-details">
                                                        <div className="kt-notification__item-title">
                                                            New order has been received
                                                        </div>
                                                        <div className="kt-notification__item-time">
                                                            2 hrs ago
                                                        </div>
                                                    </div>
                                                </a>
                                                <a href="#" className="kt-notification__item">
                                                    <div className="kt-notification__item-icon">
                                                        <i className="flaticon2-box-1 kt-font-brand"></i>
                                                    </div>
                                                    <div className="kt-notification__item-details">
                                                        <div className="kt-notification__item-title">
                                                            New customer is registered
                                                        </div>
                                                        <div className="kt-notification__item-time">
                                                            3 hrs ago
                                                        </div>
                                                    </div>
                                                </a>
                                                <a href="#" className="kt-notification__item">
                                                    <div className="kt-notification__item-icon">
                                                        <i className="flaticon2-chart2 kt-font-danger"></i>
                                                    </div>
                                                    <div className="kt-notification__item-details">
                                                        <div className="kt-notification__item-title">
                                                            Application has been approved
                                                        </div>
                                                        <div className="kt-notification__item-time">
                                                            3 hrs ago
                                                        </div>
                                                    </div>
                                                </a>
                                                <a href="#" className="kt-notification__item">
                                                    <div className="kt-notification__item-icon">
                                                        <i className="flaticon2-image-file kt-font-warning"></i>
                                                    </div>
                                                    <div className="kt-notification__item-details">
                                                        <div className="kt-notification__item-title">
                                                            New file has been uploaded
                                                        </div>
                                                        <div className="kt-notification__item-time">
                                                            5 hrs ago
                                                        </div>
                                                    </div>
                                                </a>
                                                <a href="#" className="kt-notification__item">
                                                    <div className="kt-notification__item-icon">
                                                        <i className="flaticon2-drop kt-font-info"></i>
                                                    </div>
                                                    <div className="kt-notification__item-details">
                                                        <div className="kt-notification__item-title">
                                                            New user feedback received
                                                        </div>
                                                        <div className="kt-notification__item-time">
                                                            8 hrs ago
                                                        </div>
                                                    </div>
                                                </a>
                                            </div>
                                        </div>
                                        
                                    
                                    </div>
                                </form>
                            </div>
                        </div>
            <div className="kt-header__topbar-item kt-header__topbar-item--user ml-4">
                            <div className="kt-header__topbar-wrapper" data-toggle="dropdown" data-offset="0px,0px">
                                <div className="kt-header__topbar-user">
                                    <span className="kt-header__topbar-welcome kt-hidden-mobile">Hi,</span>
                                    <span className="kt-header__topbar-username kt-hidden-mobile"> Ashutosh</span>
                                    <span className="ml-2 kt-badge kt-badge--username kt-badge--unified-success kt-badge--lg kt-badge--rounded kt-badge--bold">A</span>
                                </div>
                            </div>

                            <div className="dropdown-menu dropdown-menu-fit dropdown-menu-right dropdown-menu-anim dropdown-menu-top-unround dropdown-menu-xl">
                              
                                <div className="kt-notification">
                                    <a href="custom/apps/user/profile-1/personal-information.html" className="kt-notification__item">
                                        <div className="kt-notification__item-icon">
                                            <i className="flaticon2-calendar-3 kt-font-success"></i>
                                        </div>
                                        <div className="kt-notification__item-details">
                                            <div className="kt-notification__item-title kt-font-bold">
                                                My Profile
                                            </div>
                                            <div className="kt-notification__item-time">
                                                Account settings and more
                                            </div>
                                        </div>
                                    </a>
                                    <a href="custom/apps/user/profile-3.html" className="kt-notification__item">
                                        <div className="kt-notification__item-icon">
                                            <i className="flaticon2-mail kt-font-warning"></i>
                                        </div>
                                        <div className="kt-notification__item-details">
                                            <div className="kt-notification__item-title kt-font-bold">
                                                My Messages
                                            </div>
                                            <div className="kt-notification__item-time">
                                                Inbox and tasks
                                            </div>
                                        </div>
                                    </a>
                                    <a href="custom/apps/user/profile-2.html" className="kt-notification__item">
                                        <div className="kt-notification__item-icon">
                                            <i className="flaticon2-rocket-1 kt-font-danger"></i>
                                        </div>
                                        <div className="kt-notification__item-details">
                                            <div className="kt-notification__item-title kt-font-bold">
                                                My Activities
                                            </div>
                                            <div className="kt-notification__item-time">
                                                Logs and notifications
                                            </div>
                                        </div>
                                    </a>
                                    <a href="custom/apps/user/profile-3.html" className="kt-notification__item">
                                        <div className="kt-notification__item-icon">
                                            <i className="flaticon2-hourglass kt-font-brand"></i>
                                        </div>
                                        <div className="kt-notification__item-details">
                                            <div className="kt-notification__item-title kt-font-bold">
                                                My Tasks
                                            </div>
                                            <div className="kt-notification__item-time">
                                                latest tasks and projects
                                            </div>
                                        </div>
                                    </a>

                                    <a href="custom/apps/user/profile-1/overview.html" className="kt-notification__item">
                                        <div className="kt-notification__item-icon">
                                            <i className="flaticon2-cardiogram kt-font-warning"></i>
                                        </div>
                                        <div className="kt-notification__item-details">
                                            <div className="kt-notification__item-title kt-font-bold">
                                                Billing
                                            </div>
                                            <div className="kt-notification__item-time">
                                                billing & statements <span className="kt-badge kt-badge--danger kt-badge--inline kt-badge--pill kt-badge--rounded">2 pending</span>
                                            </div>
                                        </div>
                                    </a>
                                    <div className="kt-notification__custom kt-space-between">
                                        <a href="custom/user/login-v2.html" target="_blank" className="btn btn-label btn-label-brand btn-sm btn-bold">Sign Out</a>

                                        <a href="custom/user/login-v2.html" target="_blank" className="btn btn-clean btn-sm btn-bold">Upgrade Plan</a>
                                    </div>
                                </div>
                            </div>
                        </div>

            <button className="kt-header-mobile__toggler" id="kt_header_mobile_toggler"><span></span></button>

        </div>
    </div>
    )
}

export default HeaderMobile
