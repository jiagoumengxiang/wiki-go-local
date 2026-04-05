# Roadmap: Wiki-Go 隐藏文件读取功能

**Created:** 2025-04-05
**Phases:** 4
**Requirements Covered:** 18/18

## Overview

This roadmap breaks down the hidden files feature into 4 phases, each building on the previous one. Each phase has a clear goal with observable success criteria.

## Phase 1: 文件系统基础修改

**Goal:** 修改文件系统逻辑以支持隐藏文件夹扫描和读取

**Requirements:**
- FS-01: 系统能够扫描和读取隐藏文件夹（以 `.` 开头的文件夹）中的文件
- FS-02: 文档路径解析逻辑支持隐藏文件夹路径
- FS-03: 文档列表 API 返回隐藏文件夹中的文档

**Success Criteria:**
1. 文件系统扫描包含隐藏文件夹（如 `.config`、`.system`）
2. 路径解析器正确处理隐藏文件夹路径
3. 文档列表 API 返回包含隐藏文件夹中的文档

**Plans:**
1/1 plans complete

**Estimated Plans:** 1

---

## Phase 2: 用户界面和文档操作

**Goal:** 隐藏文件夹中的文档在侧边栏显示，并支持完整的编辑操作

**Requirements:**
- SID-01: 隐藏文件夹在侧边栏导航树中显示
- SID-02: 隐藏文件夹中的文档在侧边栏中可见
- SID-03: 隐藏文件夹的层级结构正确展示（嵌套关系）
- EDT-01: 用户可以编辑隐藏文件夹中的文档
- EDT-02: 隐藏文件夹中的文档支持保存操作
- EDT-03: 隐藏文件夹中的文档支持删除操作
- EDT-04: 隐藏文件夹中的文档支持重命名/移动操作

**Success Criteria:**
1. 隐藏文件夹在侧边栏导航树中正确显示
2. 用户可以打开并编辑隐藏文件夹中的文档
3. 用户可以保存、删除、重命名隐藏文件夹中的文档

**Estimated Plans:** 2-3

---

## Phase 3: 文件附件和版本历史

**Goal:** 支持隐藏文件夹中的文件附件和版本历史功能

**Requirements:**
- ATT-01: 用户可以上传附件到隐藏文件夹中的文档目录
- ATT-02: 隐藏文件夹中的文档可以引用附件文件
- ATT-03: 附件文件列表 API 支持隐藏文件夹路径
- VER-01: 隐藏文件夹中的文档支持版本历史记录
- VER-02: 用户可以查看隐藏文件夹中文档的历史版本
- VER-03: 用户可以恢复隐藏文件夹中文档的历史版本

**Success Criteria:**
1. 用户可以上传附件到隐藏文件夹中的文档
2. 隐藏文件夹中的文档支持版本历史
3. 用户可以查看和恢复隐藏文件夹中文档的历史版本

**Estimated Plans:** 1-2

---

## Phase 4: 搜索和权限控制

**Goal:** 确保搜索功能索引隐藏文件夹，权限系统正确应用

**Requirements:**
- SRCH-01: 搜索索引包含隐藏文件夹中的文档内容
- SRCH-02: 搜索结果包含隐藏文件夹中的文档
- SRCH-03: 搜索结果显示文档路径（包括隐藏文件夹）
- PERM-01: 隐藏文件夹的访问权限遵循现有的权限系统
- PERM-02: 私有模式下的隐藏文件夹文档需要登录才能访问
- PERM-03: 访问规则（access rules）适用于隐藏文件夹路径

**Success Criteria:**
1. 搜索功能能够找到隐藏文件夹中的文档
2. 搜索结果显示完整的文档路径（包括隐藏文件夹）
3. 隐藏文件夹的访问权限正确应用（私有模式、访问规则）

**Estimated Plans:** 1-2

---

## Progress Tracking

| Phase | Status | Plans | Progress |
|-------|--------|-------|----------|
| 1 | 1/1 | Complete   | 2026-04-04 |
| 2 | 1/1 | Complete   | 2026-04-05 |
| 3 | Pending | 1-2 | 0% |
| 4 | Pending | 1-2 | 0% |

**Total Progress:** 0%

---

## Technical Notes

### Files to Modify (Based on Architecture)

**Phase 1:**
- `internal/utils/nav.go` - Navigation tree building (likely filters out hidden directories)
- `internal/handlers/list.go` or similar - Document list API
- `internal/utils/path.go` - Path parsing and sanitization

**Phase 2:**
- `internal/handlers/editor.go` - Document edit handlers
- `internal/handlers/save.go` - Document save handlers
- Template files - Sidebar navigation display

**Phase 3:**
- `internal/handlers/attachments.go` - File upload handlers
- `internal/handlers/versions.go` - Version history handlers

**Phase 4:**
- `internal/handlers/search.go` - Search indexing and query
- `internal/auth/access.go` - Access control logic
- `internal/utils/filter.go` - Navigation filtering with access rules

### Key Decisions

- **No special configuration needed**: All hidden folders are accessible by default
- **Use existing permission system**: No special logic needed for hidden files
- **Backward compatible**: Existing functionality unchanged

---

*Roadmap created: 2025-04-05*
