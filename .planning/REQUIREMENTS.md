# Requirements: Wiki-Go 隐藏文件读取功能

**Defined:** 2025-04-05
**Core Value:** Wiki-Go 能够访问隐藏文件夹中的 Markdown 文件，实现系统配置文档的集中管理。

## v1 Requirements

Requirements for initial release. Each maps to roadmap phases.

### 文件系统读取

- [x] **FS-01**: 系统能够扫描和读取隐藏文件夹（以 `.` 开头的文件夹）中的文件
- [x] **FS-02**: 文档路径解析逻辑支持隐藏文件夹路径
- [x] **FS-03**: 文件列表 API 返回隐藏文件夹中的文档

### 侧边栏显示

- [x] **SID-01**: 隐藏文件夹在侧边栏导航树中显示
- [x] **SID-02**: 隐藏文件夹中的文档在侧边栏中可见
- [x] **SID-03**: 隐藏文件夹的层级结构正确展示（嵌套关系）

### 文档编辑

- [x] **EDT-01**: 用户可以编辑隐藏文件夹中的文档
- [x] **EDT-02**: 隐藏文件夹中的文档支持保存操作
- [x] **EDT-03**: 隐藏文件夹中的文档支持删除操作
- [x] **EDT-04**: 隐藏文件夹中的文档支持重命名/移动操作

### 文件附件

- [x] **ATT-01**: 用户可以上传附件到隐藏文件夹中的文档目录
- [x] **ATT-02**: 隐藏文件夹中的文档可以引用附件文件
- [x] **ATT-03**: 附件文件列表 API 支持隐藏文件夹路径

### 版本历史

- [x] **VER-01**: 隐藏文件夹中的文档支持版本历史记录
- [x] **VER-02**: 用户可以查看隐藏文件夹中文档的历史版本
- [x] **VER-03**: 用户可以恢复隐藏文件夹中文档的历史版本

### 搜索功能

- [x] **SRCH-01**: 搜索索引包含隐藏文件夹中的文档内容
- [x] **SRCH-02**: 搜索结果包含隐藏文件夹中的文档
- [x] **SRCH-03**: 搜索结果显示文档路径（包括隐藏文件夹）

### 权限控制

- [x] **PERM-01**: 隐藏文件夹的访问权限遵循现有的权限系统
- [x] **PERM-02**: 私有模式下的隐藏文件夹文档需要登录才能访问
- [x] **PERM-03**: 访问规则（access rules）适用于隐藏文件夹路径

## v2 Requirements

Deferred to future release. Tracked but not in current roadmap.

### 高级功能

- **ADV-01**: 配置选项允许控制哪些隐藏文件夹可访问（白名单/黑名单）
- **ADV-02**: 管理界面显示隐藏文件夹的特殊标记
- **ADV-03**: 隐藏文件夹的批量管理功能

## Out of Scope

Explicitly excluded. Documented to prevent scope creep.

| Feature | Reason |
|---------|--------|
| 隐藏文件访问 | 只关注文件夹（如 `.config`），不涉及隐藏文件（如 `.gitignore`） |
| 系统级文件 | 仅限于 `data/documents/` 目录下的隐藏文件夹，不访问系统路径 |
| 特殊权限系统 | 使用现有的权限系统，不创建专门的隐藏文件权限逻辑 |
| 性能优化 | 在当前实现中不进行专门的性能优化 |

## Traceability

Which phases cover which requirements. Updated during roadmap creation.

| Requirement | Phase | Status |
|-------------|-------|--------|
| FS-01 | Phase 1 | Complete |
| FS-02 | Phase 1 | Complete |
| FS-03 | Phase 1 | Complete |
| SID-01 | Phase 2 | Complete |
| SID-02 | Phase 2 | Complete |
| SID-03 | Phase 2 | Complete |
| EDT-01 | Phase 2 | Complete |
| EDT-02 | Phase 2 | Complete |
| EDT-03 | Phase 2 | Complete |
| EDT-04 | Phase 2 | Complete |
| ATT-01 | Phase 3 | Complete |
| ATT-02 | Phase 3 | Complete |
| ATT-03 | Phase 3 | Complete |
| VER-01 | Phase 3 | Complete |
| VER-02 | Phase 3 | Complete |
| VER-03 | Phase 3 | Complete |
| SRCH-01 | Phase 4 | Complete |
| SRCH-02 | Phase 4 | Complete |
| SRCH-03 | Phase 4 | Complete |
| PERM-01 | Phase 4 | Complete |
| PERM-02 | Phase 4 | Complete |
| PERM-03 | Phase 4 | Complete |

**Coverage:**
- v1 requirements: 18 total
- Mapped to phases: 18
- Unmapped: 0 ✓

---
*Requirements defined: 2025-04-05*
*Last updated: 2025-04-05 after initial definition*
