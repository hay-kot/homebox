<div align="center">
  <img src="/docs/docs/assets/img/lilbox.svg" height="200"/>
</div>

<h1 align="center" style="margin-top: -10px"> HomeBox </h1>
<p align="center" style="width: 100;">
   <a href="https://hay-kot.github.io/homebox/">Docs</a>
   |
   <a href="https://homebox.fly.dev">Demo</a>
   |
   <a href="https://discord.gg/tuncmNrE4z">Discord</a>
</p>

## Quick Start

```yml
version: "3.4"
 services:
   homebox:
     image: ghcr.io/hay-kot/homebox:nightly
     container_name: homebox
     restart: always
     volumes:
       - homebox-data:/data/
     ports:
       - 3100:7745

volumes:
   homebox-data:
     driver: local
```

## MVP Todo

- [ ] Asset Attachments for Items
- [ ] Db Migrations
  - [ ] How To
- [x] Documentation
  - [x] Docker Compose
  - [x] Config Options
- [x] Locations
  - [x] Create
  - [x] Update
  - [x] Delete
- [x] Labels
  - [x] Create
  - [x] Update
  - [x] Delete
- [x] Items CRUD
  - [x] Create
  - [x] Update
  - [x] Delete
- [x] Fields To Add
  - [x] Quantity
  - [x] Insured (bool)
- [x] Bulk Import via CSV
  - [x] Initial
  - [x] Add Warranty Columns
  - [x] All Fields
  - [x] Documentations
- [x] Release Flow
  - [x] CI/CD Docker Builds w/ Multi-arch
  - [x] Auto Fly.io Deploy for Nightly
  - [x] Deploy Docs
- [x] Repo House Keeping
  - [x] Add License
  - [x] Issues Template
  - [x] PR Templates
  - [x] Security Policy
  - [x] Feature Request Template
- [x] Embedded Version Info
  - [x] Version Number
  - [x] Git Hash
- [x] Setup Docker Volumes in Dockerfile
- [x] Warranty Information
  - [x] Option for Lifetime Warranty or Warranty Period

## All Todo's

- [ ] Dev Container for Development
- [ ] User Invitation Links to Join Group
- [ ] Maintenance Logs
  - [ ] Schedule Future Maintenance
  - [ ] Email on Maintenance Due
  - [ ] QR Code Stickers to Scan to enter a Maintenance Task
- [ ] Export CSV (With IDs)
- [ ] User Profile
  - [ ] Adjust Theme (Daisy UI)
  - [ ] Delete Profile
  - [ ] Send User Invites
  - [ ] Set Currency
  - [ ] Change Password
- [ ] Admin Page
  - [ ] Instance Statistics
  - [ ] User Management
    - [ ] Delete User
    - [ ] Reset Password
- [ ] Expose Swagger API Documentation
  - [ ] Dynamic Port / Host Settings

## Credits

- Logo by [@lakotelman](https://github.com/lakotelman)
