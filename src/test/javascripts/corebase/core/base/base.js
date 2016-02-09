(function() {
    'use strict';

    angular
        .module('app.core')
        .controller('core', Core);

    Core.$inject = [
        '$rootScope', '$translate', 'config', 'common'
    ];

    /*@ngInject*/
    function Core($rootScope, $translate, config, common) {
        /* jshint validthis:true */
        var main = this;
        var logger = common.logger;

        //variables
        main.langAvailable = [];
        main.langDefault = 'eng';

        // methods
        main.onLangChange = onLangChange;

        $rootScope.main = main;

        activate();
        // activation

        // Update user when has update into $auth.
        $rootScope.$on('auth:updated', function () {
            activate();
        });

        //onLangChange
        function activate() {
            main.langAvailable = config.language.available;
            main.langDefault   = $translate.use() || config.language.default;

            logger.debug('Core loaded!', null);
        }

        function onLangChange(lang) {
            if (!!main.langAvailable[lang]) {
                main.langDefault = lang;
                $translate.use(lang);
            }
        }
    }
})();
