## postman scripts里面post-response
```
const jsonData = pm.response.json();
const captchaDataUri = jsonData.data.captcha;

const template = `
  <html>
    <body style="text-align:center; padding-top:20px;">
      <img src="${captchaDataUri}" alt="验证码" style="max-width:90%;"/>
    </body>
  </html>
`;
pm.environment.set("captchaID", jsonData.data.captchaID);
pm.visualizer.set(template);
```