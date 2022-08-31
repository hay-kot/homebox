import { Requests } from '../../lib/requests';
//  <
// 	TGetResult,
// 	TPostData,
// 	TPostResult,
// 	TPutData = TPostData,
// 	TPutResult = TPostResult,
// 	TDeleteResult = void
// >

export class BaseAPI {
  http: Requests;

  constructor(requests: Requests) {
    this.http = requests;
  }
}
