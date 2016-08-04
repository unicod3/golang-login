<<< template "../base.tpl" . >>>

<<< define "css" >>>
<link rel="stylesheet" href='/static/css/custom.css' />
<<< end >>>

<<< define "content" >>>
<div class="l-home"><div class="l-home container">
  <div class="l-home__user">
    <h1>Hello <<<.Username>>></h1>
    <p>Company: <strong><<<.Company>>></strong></p>
    <p>Categories:  <strong>Category 1 </strong><strong>Category 2 </strong><strong>Category 3 </strong></p>
    <p>Address 1: <strong>California</strong></p>
    <p>Address 2: <strong>Elm Street</strong></p>
    <p>Phone: <strong>+1 5599999999</strong></p>

    <footer>
      <a href="<<<urlfor "UserAuthController.Logout">>>" class="button button--primary button--filled">Logout</a>
      <button class="button button--primary button--filled">Edit</button>
    </footer>
  </div>
</div>
</div>
<<< end >>>

<<< define "js" >>>
<script src="/static/js/custom.js"></script>
<<< end >>>
