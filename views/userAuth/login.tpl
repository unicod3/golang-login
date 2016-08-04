<<< template "../base.tpl" . >>>

<<< define "css" >>>
<link rel="stylesheet" href='/static/css/custom.css' />
<style>
body{
    background:#7A7B7C;
    padding-top:1.5em;
}
</style>
<<< end >>>


<<< define "content" >>>
<div class="l-auth">
<div class="l-auth__wrapper">
      <button class="l-auth__close"></button>

      <router-outlet></router-outlet><div class="l-auth--login">
      <div class="l-auth__body">
  <header class="l-auth__header">
    <h1 class="l-auth__title">Log In</h1>
  </header>

  <button class="l-auth__provider l-auth__provider--linkedin button button--filled">
    <span class="l-auth__provider-icon"></span>
    Log In with Linked In
  </button>

<form role="form" class="form form--default" method="POST" action='<<<urlfor "UserAuthController.Login">>>'>
  <div class="l-auth__divisor">Or</div>
<<< .xsrfdata >>>
    <div class="l-auth__email form__group">
      <label class="form__label">E-mail</label>
      <input type="email" class="form__input ng-untouched ng-pristine ng-invalid" name="email" placeholder="" required="" ng-reflect-name="email">
    </div>
    <div class="l-auth__password form__group">
      <label class="form__label">Password</label>
      <input type="password" class="form__input ng-untouched ng-pristine ng-invalid" name="password" placeholder="" required="" ng-reflect-name="password">
    </div>
    
    <div class="form__group text-danger">
        <<<.flash.error>>>
    </div>
    
</div>
<footer class="l-auth__footer">
  <a class="l-auth__forgot-password" ng-reflect-router-link="../password" ng-reflect-href="/welcome/auth/password" href='<<<urlfor "UserAuthController.Password">>>'>Forgot password?</a>
  <button type="submit" class="l-auth__continue button button--filled button--primary button--lg">Log In</button>
  <p class="l-auth__existing">Do not have an account? <a ng-reflect-router-link="../register" ng-reflect-href="/welcome/auth/register" href='<<<urlfor "UserAuthController.Register">>>'>Sign Up!</a></p>
</footer>
  </form>
</div>
    </div>
<<< end >>>


<<< define "js" >>>
<script src="/static/js/custom.js"></script>
<<< end >>>


