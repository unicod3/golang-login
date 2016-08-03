<<< template "../base.tpl" . >>>

<<< define "css" >>>
<link rel="stylesheet" href='/static/css/custom.css' />
<<< end >>>


<<< define "content" >>>
<div class="container">
   Email :  <<< .userEmail >>><br>
   username:  <<< .username >>><br>
   password :  <<< .password >>>
</div>
<<< end >>>


<<< define "js" >>>
<script src="/static/js/custom.js"></script>
<<< end >>>


