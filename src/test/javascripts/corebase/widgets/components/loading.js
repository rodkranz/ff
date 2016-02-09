(function () {
    'use strict';

    angular.module('app.widgets')
        .directive('loadingContainer', LoadingContainer);

    LoadingContainer.$inject = ['logger'];

    function LoadingContainer(logger)
    {
        var services = {
            restrict:   'A',
            link:       getLink,
            scope:      false
        };

        return services;
        ////////////////

        function getLink(scope, element, attrs)
        {
            var loadingLayer = angular.element("<div class=\"loading\"><div class=\"loading-spinner\"><span class=\"fa fa-refresh spin\"></span></div></div>");
            element.append(loadingLayer);
            element.addClass("loading-container");
            scope.$watch(attrs.loadingContainer, function (value) {
                loadingLayer.toggleClass("ng-hide", !value);
            });
        }
    }
})();
