<html>
    <head>
     <title>Kayıt</title>
    <link rel="stylesheet" href="/common.css">
    </head>
    <body>
        <h1>Merhaba</h1>
    <p>Lütfen Kullanıcı Adı ve Kullanıcı Soyadı'nı giriniz. <br/>
   
  
    
    
        <form action="/login" method="post">
            Kullanıcı Adı:<br/><input type="text" name="username"><br/>
          <p> Kullanıcı Soyadı:<br/><input type="text" name="password"><br/><br/>
           
               <button onclick="window.location.href = '/login';">Kayıt</button>
       <input type="button" onclick="window.location.href = '/';" value="Geri"/><br/><br/><br/>
                 Teşekkürler...</p>
        </form>
    </body>
</html>