<<< template "../base.tpl" . >>>

<<< define "css" >>>
<link rel="stylesheet" href='/static/css/custom.css' />
<<< end >>>

<<< define "content" >>>
  <div class="l-welcome" ng-view>
  <div class="l-welcome__wrapper">
  <div class="l-welcome__container">
    <h1 class="l-welcome__brand">Go Test</h1>
    <a  href='<<<urlfor "UserAuthController.Login">>>' class="button button--filled button--primary">Login</a>
  
    <a [routerLink]="['./auth/register']" class="button button--filled button--primary">Sign Up</a>


  </div>
<<< end >>>

<<< define "js" >>>
<script src="/static/js/custom.js"></script>
<<< end >>>
