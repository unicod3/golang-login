{{ template "../base.tpl" . }}

{{ define "css" }}
<link rel="stylesheet" href='/static/css/custom.css' />
{{ end }}


{{ define "content" }}
<div class="container">
    <div class="row vertical-offset-75">
    	<div class="col-md-6 col-md-offset-3">
    		<div class="panel panel-default">
			  	<div class="panel-heading text-center">
			    	<h3 class="panel-title"><strong>Login</strong></h3>
			 	</div> 

			  	<div class="panel-body">
			    	<form role="form" class="form-horizontal" method="POST" action='{{urlfor "UserAuthController.Login"}}'>
                      {{ .xsrfdata }}

                      <div class="form-group">
                        <label for="inputEmail" class="col-sm-3 control-label">Username</label>
                        <div class="col-sm-9">
                          <input class="form-control" placeholder="Username" name="Username" value="" type="text" required id="inputEmail" />
                        </div>
                      </div>
                      <div class="form-group">
                        <label for="inputPassword" class="col-sm-3 control-label">Password</label>
                        <div class="col-sm-9">
			    		  <input class="form-control" placeholder="Password" name="Password" type="password" required id="inputPassword"  />
                        </div>
                      </div>
                      <div class="form-group">
                        <div class="col-sm-2 pull-right">
			    		    <input class="btn btn-lg btn-success" type="submit" value="Submit"> 
                        </div>
                      </div>
                      <div class="form-group">
                        <div class="col-sm-3 pull-right">
                            <a href="{{urlfor "UserAuthController.PasswordReset"}}"> 
                                forgot password »
                            </a>
                        </div>
                      </div>
                    </form>
			    </div>

                <div class="panel-footer text-center clearfix"> <a href='{{urlfor "UserAuthController.Register"}}'>Register »</a></div>

			</div>
		</div>
	</div>
</div>
{{ end }}


{{ define "js" }}
<script src="/static/js/custom.js"></script>
{{ end }}


