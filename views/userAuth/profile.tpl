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
    <h1 class="l-auth__title">Fill Your Profile</h1>
  </header>

<form role="form" class="form form--default" method="POST" action='<<<urlfor "UserAuthController.Profile">>>'>
<<< .xsrfdata >>>
    <div class="l-auth__group l-auth__group--firstName">
    <label class="form__label">First Name</label>
    <input type="text" class="form__input ng-pristine ng-invalid ng-touched" name="firstName" value='<<< .firstName >>>' placeholder="" required="" ng-reflect-name="firstName">
  </div>
  <div class="l-auth__group l-auth__group--companyName">
    <label class="form__label">Last Name</label>
    <input type="text" class="form__input ng-untouched ng-pristine ng-invalid" name="lastName"  value='<<< .lastName >>>' placeholder="" required="" ng-reflect-name="lastName">
  </div>

  <div class="l-auth__group l-auth__group--company">
    <label class="form__label">Company Name</label>
    <input type="text" class="form__input ng-dirty ng-valid ng-touched" name="company"  value='<<< .company >>>' placeholder="" required="" ng-reflect-name="company" ng-reflect-model="www">
  </div>

  <div class="l-auth__group l-auth__group--categories">
    <label class="form__label">Categories</label>
    <input type="text" class="form__input ng-untouched ng-pristine ng-invalid" name="categories" value='<<< .categories >>>'  required="" ng-reflect-model="" ng-reflect-name="categories" ng-reflect-typeahead="Alabama,Alaska,Arizona,Arkansas,California,Colorado,Connecticut,Delaware,Florida,Georgia,Hawaii,Idaho,Illinois,Indiana,Iowa,Kansas,Kentucky,Louisiana,Maine,Maryland,Massachusetts,Michigan,Minnesota,Mississippi,Missouri,Montana,Nebraska,Nevada,New Hampshire,New Jersey,New Mexico,New York,North Dakota,North Carolina,Ohio,Oklahoma,Oregon,Pennsylvania,Rhode Island,South Carolina,South Dakota,Tennessee,Texas,Utah,Vermont,Virginia,Washington,West Virginia,Wisconsin,Wyoming">
  </div>

  
  <div class="l-auth__categories">
    <!--template bindings={
  "ng-reflect-ng-for-of": ""
}-->
  </div>

  <div class="l-auth__group l-auth__group--address-1">
    <label class="form__label">Address</label>
    <input type="text" class="form__input ng-untouched ng-pristine ng-invalid" name="address1"  value='<<< .address1 >>>' placeholder="Street Address &amp; Number" required="" ng-reflect-name="address1">
  </div>

  <div class="l-auth__group l-auth__group--address-2">
    <label class="form__label">Address</label>
    <input type="text" class="form__input ng-untouched ng-pristine ng-invalid" name="address2" value='<<< .address2 >>>'  placeholder="City State &amp; Zip" required="" ng-reflect-name="address2">
  </div>

  <div class="l-auth__group l-auth__group--phone">
    <label class="form__label">Phone</label>
    <input type="text" class="form__input ng-untouched ng-pristine ng-invalid" name="phone" value='<<< .phone >>>'  placeholder="Phone Number" required="" ng-reflect-name="phone">
  </div>
    
    <div class="form__group text-danger">
        <<<.flash.error>>>
    </div>
    
</div>
<footer class="l-auth__footer">
  <button type="submit" class="l-auth__continue button button--filled button--primary button--lg">Continue</button>
</footer>
  </form>
</div>
    </div>
<<< end >>>


<<< define "js" >>>
<script src="/static/js/custom.js"></script>
<<< end >>>


