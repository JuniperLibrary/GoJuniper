# OpenCode Skills 使用教程：从入门到实战

## 一、什么是OpenCode Skills？

OpenCode Skills 是可安装的模块化扩展，为OpenCode AI编程助手提供特定领域的专业知识和能力。通过安装不同的Skills，可以让AI在特定领域（如UI/UX设计、视频生成、特定框架开发等）提供更专业的帮助。

## 二、Skills安装方法

### 1. 安装命令
```bash
npx skills add <owner/repo>
```
例如：
```bash
npx skills add nextlevelbuilder/ui-ux-pro-max-skill
```

### 2. 安装过程
1. 运行安装命令后，系统会提示选择安装范围：
   - **当前目录**：仅在当前项目中可用（推荐用于项目开发）
   - **全局安装**：在所有项目中可用（适用于通用工具）

2. 安装过程中可以选择支持的开发工具（如OpenCode、VS Code等）

3. 安装完成后，会显示安装目录和支持的开发工具信息

## 三、在OpenCode中使用Skills

### 1. 调用已安装的Skills
- 打开OpenCode终端
- 输入 `/` 后跟skill名称的首字母或相关关键字
  - 例如：输入 `/ui` 可以看到已安装的UI相关skills
  - 例如：输入 `/remotion` 可以看到已安装的Remotion相关skills

### 2. 使用流程
1. 调用skill后，AI会显示可用的skill列表
2. 选择需要使用的skill
3. 提供自然语言描述你的需求
4. AI会自动调用skill的知识库来理解并执行你的请求
5. 生成相应的代码、设计或其他输出

## 四、实战案例演示

### 案例1：使用ui-ux-pro-max skill创建着陆页
1. 安装skill：`npx skills add nextlevelbuilder/ui-ux-pro-max-skill`
2. 在OpenCode中输入 `/ui` 并选择ui-ux-pro-max skill
3. 提供需求：`为宠物美容服务搭建一个着陆页，风格活泼亲和，并设置预约类行动召唤按钮`
4. AI会：
   - 调用skill的设计知识库
   - 生成目录结构和文件
   - 提供符合活泼亲和风格的UI设计
   - 包含预约按钮等交互元素

### 案例2：使用remotion-best-practices skill生成视频
1. 安装skill：`npx skills add remotion-dev/skills`
2. 在OpenCode中输入 `/remotion` 并选择remotion-best-practices skill
3. 提供需求：`生成一个 Hello Runoob！的演示视频`
5. AI会：
   - 应用Remotion最佳实践
   - 创建视频项目结构
   - 生成包含文字动画的视频代码
   - 提供媒体导入、序列组织等专业指导

## 五、技巧和最佳实践

### 1. 选择合适的安装范围
- 项目特定skills：选择当前目录安装，避免污染全局环境
- 通用工具skills：可以考虑全局安装

### 2. 有效使用skills的方法
- 具体描述需求：越具体，AI调用skill的效果越好
- 利用skill的专业领域：每个skill都有其特长，发挥其优势
- 迭代优化：先生成基础版本，再根据反馈改进

### 3. 常见问题排查
- 如果看不到已安装的skill：确认安装成功且在正确的目录中
- 如果skill不起作用：尝试重新安装或检查skill是否与当前OpenCode版本兼容
- 性能问题：有些skill可能需要较大的上下文，确保有足够的资源

## 六、进阶应用

### 1. 组合使用多个skills
复杂项目可以组合使用多个skills：
- 先用ui-ux-pro-max设计界面
- 用remotion-best-practices创建演示视频
- 用特定框架skills实现功能

### 2. 自定义skills
随着经验积累，可以学习创建自己的skills来解决特定问题或共享团队最佳实践。

## 七、总结

OpenCode Skills显著增强了AI编程助手的能力，使其能够在特定领域提供专业水平的帮助。通过正确安装和使用skills，开发者可以：

1. 提高开发效率
2. 获得专业领域的指导
3. 减少在特定技术上的学习曲线
4. 聚焦于业务逻辑而非重复性任务

掌握skills的使用是成为高效OpenCode用户的重要一步，从简单的页面设计到复杂的视频生成，skills都能提供有价值的专业支持。