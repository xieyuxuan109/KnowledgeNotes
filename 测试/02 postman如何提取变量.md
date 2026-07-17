## 提取环境变量代码
```
const jsonData = pm.response.json();
pm.environment.set("captchaID", jsonData.data.captchaID);
```