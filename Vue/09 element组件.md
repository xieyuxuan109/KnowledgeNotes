## el-radio diabled表示是否禁用 border表示是否加上边框 size表示大小 单选框
```
   <el-avatar shape = "circle" :src = "url" />
   <el-radio-group v-model="sex">
   <el-radio value='1'>男</el-radio>
   <el-radio value='0'>女</el-radio>
   </el-radio-group>
    </div>
```
## 复选框 el-checkbox el-checkbox-button按钮样式
## el-input claerable type="password" show-password prefix-icon suffix-icon maxlength minlength
```
type属性 text 文本框 textarea 文本区域 password 密码框 button 按钮 file 文件
```
## 常用事件 blur事件 当选择器的输入框失去焦点时触发
```
   <el-input v-model="inputText" style="width: 240px;" placeholder="请输入文本" @blur="onBlur"></el-input>
    </div>
  </div>

</template>
<script setup>
// 确保在 <script setup> 中导入了图标组件
import { Plus, EditPen } from '@element-plus/icons-vue'
import {reactive,toRefs,ref} from 'vue'
import {ElMessage } from 'element-plus'
const state = reactive({
  url:'https://cube.elemecdn.com/3/7c/3ea6beec64369c2642b92c6726f1epng.png'
})
const sex=ref("1")
const {url}=toRefs(state)
const inputText=ref('')
const onBlur = ()=>{
  ElMessage.success("调用成功！！！")
}
```
## focus时件当有焦点时触发
## change时件仅当modelValue高便是 当输入框失去焦点或用户按下enter键触发
## input函数 当input值改变的时候触发
## clear事件 再点击clearable属性生成的清空按钮时触发的事件
## el-input插槽用法 prefix suffix prepend append四种属性
## 带推荐列表的输入框 el-autocomplete
```
<el-autocomplete v-model="searchValue" :fetch-suggestions="querySearch" clearable placeholder="请输入查询内容" @selete="handleSelect"></el-autocomplete>
```
## debounce延时显示
## fetch-suggestions属性 获取输入框建议的函数
## placement设置菜单弹出的位置 top在顶部弹出
## highlight-first-item是否选择高亮远程搜索结果第一项
## el-input-number数字输入框 min用来设置允许输入的最小值 max用来设置允许输入的最大值 步长属性step和step-strictly precision表示精度，保留几位小数 size表示大小尺寸 control-postion控制按钮位置
## el-select组件与el-option搭配使用 multiple表示多选 filterable表示启用搜索功能 visible-change事件下拉框出现或隐藏时出现
## 级联选择器el-cascader props.expandTrigger可以定义展开子菜单的触发方式