//---------------------------------------------
// Application
//---------------------------------------------

/**
 * All module that need to be loaded.
 */
//= require ./admin/app.module


/**
 * Core Base
 **/
//= require corebase
/**
 * End Core Base
 **/


/**
 * Application
 **/
// Configurations
//= require ./admin/core/configuration.js

// Data Repositories
//= require_tree ./admin/data/repositories

// Layout
//= require ./admin/layout/layout.module
//= require ./admin/layout/layout.controller
//= require ./admin/layout/layout.router

// Layout - SideBar
//= require ./admin/layout/sidebar/sidebar
//= require ./admin/layout/sidebar/sidebar.controller

// Layout - Navbar
//= require ./admin/layout/navbar/navbar
//= require ./admin/layout/navbar/navbar.controller

// Layout - Grid
//= require ./admin/layout/grid/grid
//= require ./admin/layout/grid/grid.controller

// Layout - Page
//= require admin/layout/empty/empty.js
//= require admin/layout/empty/empty.controller.js

// Layout - Footer
//= require ./admin/layout/footer/footer
//= require ./admin/layout/footer/footer.controller

// Widgets
//= require_tree ./admin/widgets/
/**
 * Core Base
 **/

/**
 * Pages
 */
// FourOhFour
//= require ./admin/fourOhFour/fourOhFour
//= require ./admin/fourOhFour/fourOhFour.module
//= require ./admin/fourOhFour/fourOhFour.route
//= require ./admin/fourOhFour/fourOhFour.controller

// Login
//= require ./admin/login/login
//= require ./admin/login/login.module
//= require ./admin/login/login.route
//= require ./admin/login/login.controller

// Dashboard
//= require ./admin/dashboard/dashboard
//= require ./admin/dashboard/dashboard.module
//= require ./admin/dashboard/dashboard.route
//= require ./admin/dashboard/dashboard.controller

// Daily Post
//= require ./admin/dailyPost/dailyPostForm
//= require ./admin/dailyPost/dailyPostList
//= require ./admin/dailyPost/dailyPost.module
//= require ./admin/dailyPost/dailyPost.route
//= require ./admin/dailyPost/dailyPostForm.controller
//= require ./admin/dailyPost/dailyPostList.controller

// School
//= require ./admin/school/schoolForm
//= require ./admin/school/schoolList
//= require ./admin/school/school.module
//= require ./admin/school/school.route
//= require ./admin/school/schoolForm.controller
//= require ./admin/school/schoolList.controller

// School - Address
//=require ./admin/school/address/addressForm
//=require ./admin/school/address/addressForm.controller

// School - Certificate
//=require ./admin/school/certificate/certificateForm
//=require ./admin/school/certificate/certificateForm.controller
//=require ./admin/school/certificate/certificateList
//=require ./admin/school/certificate/certificateList.controller

// School - Contact
//=require ./admin/school/contact/contactForm
//=require ./admin/school/contact/contactForm.controller

// School - Gallery
//=require ./admin/school/gallery/galleryForm
//=require ./admin/school/gallery/galleryForm.controller
//=require ./admin/school/gallery/galleryList
//=require ./admin/school/gallery/galleryList.controller

// School - Nationality
//=require ./admin/school/nationality/nationalityForm
//=require ./admin/school/nationality/nationalityForm.controller
//=require ./admin/school/nationality/nationalityList
//=require ./admin/school/nationality/nationalityList.controller

/**
 * End Pages
 **/
