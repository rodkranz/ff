(function () {
    'use strict';

    angular
        .module('app.widgets')
        .directive('loader', Loader);

    Loader.$inject = [];

    /* @ngInject */
    function Loader(){

        var service = {
            restrict: "EAC",
            template: getTemplate()
        };

        return service;

        function getTemplate() {
            return '' +
                '<div class="windows8"> ' +
                '   <div class="wBall" id="wBall_1"> ' +
                '       <div class="wInnerBall"> ' +
                '       </div> ' +
                '   </div> ' +
                '   <div class="wBall" id="wBall_2"> ' +
                '       <div class="wInnerBall"> ' +
                '       </div> ' +
                '   </div> ' +
                '   <div class="wBall" id="wBall_3"> ' +
                '       <div class="wInnerBall"> ' +
                '       </div> ' +
                '   </div> ' +
                '   <div class="wBall" id="wBall_4"> ' +
                '       <div class="wInnerBall"> ' +
                '       </div> ' +
                '   </div> ' +
                '   <div class="wBall" id="wBall_5"> ' +
                '       <div class="wInnerBall"> ' +
                '       </div> ' +
                '   </div> ' +
                '</div> ';
        }
    }
})();