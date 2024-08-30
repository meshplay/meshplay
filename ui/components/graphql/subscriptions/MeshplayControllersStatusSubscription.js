import { graphql, requestSubscription } from 'react-relay';
import { createRelayEnvironment } from '../../../lib/relayEnvironment';

const meshplayControllersStatusSubscription = graphql`
  subscription MeshplayControllersStatusSubscription($connectionIDs: [String!]) {
    subscribeMeshplayControllersStatus(connectionIDs: $connectionIDs) {
      connectionID
      controller
      status
      version
    }
  }
`;

export default function subscribeMeshplayControllersStatus(dataCB, variables) {
  const environment = createRelayEnvironment({});
  return requestSubscription(environment, {
    subscription: meshplayControllersStatusSubscription,
    variables: { connectionIDs: variables },
    onNext: dataCB,
    onError: (error) => console.log(`An error occured:`, error),
  });
}
