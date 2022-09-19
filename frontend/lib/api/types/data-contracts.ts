/* post-processed by ./scripts/process-types.py */
/* eslint-disable */
/* tslint:disable */
/*
 * ---------------------------------------------------------------
 * ## THIS FILE WAS GENERATED VIA SWAGGER-TYPESCRIPT-API        ##
 * ##                                                           ##
 * ## AUTHOR: acacode                                           ##
 * ## SOURCE: https://github.com/acacode/swagger-typescript-api ##
 * ---------------------------------------------------------------
 */

export interface ServerResult {
  details: any;
  error: boolean;
  item: any;
  message: string;
}

export interface ServerResults {
  items: any;
}

export interface ApiSummary {
  build: Build;
  health: boolean;
  message: string;
  title: string;
  versions: string[];
}

export interface Build {
  buildTime: string;
  commit: string;
  version: string;
}

export interface DocumentOut {
  id: string;
  path: string;
  title: string;
}

export interface ItemAttachment {
  createdAt: Date;
  document: DocumentOut;
  id: string;
  type: string;
  updatedAt: Date;
}

export interface ItemCreate {
  description: string;
  labelIds: string[];

  /** Edges */
  locationId: string;
  name: string;
}

export interface ItemOut {
  attachments: ItemAttachment[];
  createdAt: Date;
  description: string;
  id: string;
  insured: boolean;
  labels: LabelSummary[];

  /** Warranty */
  lifetimeWarranty: boolean;

  /** Edges */
  location: LocationSummary;
  manufacturer: string;
  modelNumber: string;
  name: string;

  /** Extras */
  notes: string;
  purchaseFrom: string;

  /** @example 0 */
  purchasePrice: string;

  /** Purchase */
  purchaseTime: Date;
  quantity: number;

  /** Identifications */
  serialNumber: string;
  soldNotes: string;

  /** @example 0 */
  soldPrice: string;

  /** Sold */
  soldTime: Date;
  soldTo: string;
  updatedAt: Date;
  warrantyDetails: string;
  warrantyExpires: Date;
}

export interface ItemSummary {
  createdAt: Date;
  description: string;
  id: string;
  insured: boolean;
  labels: LabelSummary[];

  /** Warranty */
  lifetimeWarranty: boolean;

  /** Edges */
  location: LocationSummary;
  manufacturer: string;
  modelNumber: string;
  name: string;

  /** Extras */
  notes: string;
  purchaseFrom: string;

  /** @example 0 */
  purchasePrice: string;

  /** Purchase */
  purchaseTime: Date;
  quantity: number;

  /** Identifications */
  serialNumber: string;
  soldNotes: string;

  /** @example 0 */
  soldPrice: string;

  /** Sold */
  soldTime: Date;
  soldTo: string;
  updatedAt: Date;
  warrantyDetails: string;
  warrantyExpires: Date;
}

export interface ItemUpdate {
  description: string;
  id: string;
  insured: boolean;
  labelIds: string[];

  /** Warranty */
  lifetimeWarranty: boolean;

  /** Edges */
  locationId: string;
  manufacturer: string;
  modelNumber: string;
  name: string;

  /** Extras */
  notes: string;
  purchaseFrom: string;

  /** @example 0 */
  purchasePrice: string;

  /** Purchase */
  purchaseTime: Date;
  quantity: number;

  /** Identifications */
  serialNumber: string;
  soldNotes: string;

  /** @example 0 */
  soldPrice: string;

  /** Sold */
  soldTime: Date;
  soldTo: string;
  warrantyDetails: string;
  warrantyExpires: Date;
}

export interface LabelCreate {
  color: string;
  description: string;
  name: string;
}

export interface LabelOut {
  createdAt: Date;
  description: string;
  id: string;
  items: ItemSummary[];
  name: string;
  updatedAt: Date;
}

export interface LabelSummary {
  createdAt: Date;
  description: string;
  id: string;
  name: string;
  updatedAt: Date;
}

export interface LocationCount {
  createdAt: Date;
  description: string;
  id: string;
  itemCount: number;
  name: string;
  updatedAt: Date;
}

export interface LocationCreate {
  description: string;
  name: string;
}

export interface LocationOut {
  createdAt: Date;
  description: string;
  id: string;
  items: ItemSummary[];
  name: string;
  updatedAt: Date;
}

export interface LocationSummary {
  createdAt: Date;
  description: string;
  id: string;
  name: string;
  updatedAt: Date;
}

export interface TokenResponse {
  expiresAt: string;
  token: string;
}

export interface UserIn {
  email: string;
  name: string;
  password: string;
}

export interface UserOut {
  email: string;
  groupId: string;
  groupName: string;
  id: string;
  isSuperuser: boolean;
  name: string;
}

export interface UserRegistration {
  groupName: string;
  user: UserIn;
}

export interface UserUpdate {
  email: string;
  name: string;
}
