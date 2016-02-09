(function () {
    'use strict';

    angular.module('app', [
          'app.core'         // Core of application
        , 'app.data'         // needs core
        , 'app.widgets'      // needs core

        , 'app.layout'       // layout's base first load

        /************************************
         * All that page need the layout base
         ************************************/
        , 'app.fourOhFour'
        , 'app.login'
        , 'app.dashboard'
        , 'app.dailyPost'
        , 'app.school'

    ]);

})();
