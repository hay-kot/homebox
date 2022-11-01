import { BaseAPI, route } from "../base";
import { Group, GroupInvitation, GroupInvitationCreate, GroupStatistics, GroupUpdate } from "../types/data-contracts";

export class GroupApi extends BaseAPI {
  createInvitation(data: GroupInvitationCreate) {
    return this.http.post<GroupInvitationCreate, GroupInvitation>({
      url: route("/groups/invitations"),
      body: data,
    });
  }

  update(data: GroupUpdate) {
    return this.http.put<GroupUpdate, Group>({
      url: route("/groups"),
      body: data,
    });
  }

  get() {
    return this.http.get<Group>({
      url: route("/groups"),
    });
  }

  statistics() {
    return this.http.get<GroupStatistics>({
      url: route("/groups/statistics"),
    });
  }
}
