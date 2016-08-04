<<< template "../base.tpl" . >>>

<<< define "css" >>>
<link rel="stylesheet" href='/static/css/custom.css' />
<<< end >>>

<<< define "content" >>>
<div class="l-home"><div class="l-home container">
  <div class="l-home__user">
    <h1>Hello <<<.firstName>>> <<<.lastName>>></h1>
    <p>Email: <strong><<<.email>>></strong></p>
    <p>Company: <strong> <<< .company >>> </strong></p>
    <p>Categories:  <strong> <<< .categories >>>  </strong></p>
    <p>Address 1: <strong><<< .address1 >>></strong></p>
    <p>Address 2: <strong><<< .address2 >>></strong></p>
    <p>Phone: <strong><<< .phone >>></strong></p>

    <footer>
      <a href='<<<urlfor "UserAuthController.Logout">>>' class="button button--primary button--filled">Logout</a>
      <a href='<<<urlfor "UserAuthController.Profile">>>' class="button button--primary button--filled">Edit</a>
    </footer>
  </div>
</div>
</div>
<<< end >>>

<<< define "js" >>>
<script src="/static/js/custom.js"></script>
<<< end >>>
