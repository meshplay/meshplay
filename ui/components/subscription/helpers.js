import subscribeMeshplayControllersStatus from '../graphql/subscriptions/MeshplayControllersStatusSubscription';
import { isMeshplayControllerStateSubscriptionDataUpdated } from './comparatorFns';
import { mergeMeshplayController } from './mergeFns';

// export const MESHSYNC_EVENT_SUBSCRIPTION = 'MESHSYNC_EVENT_SUBSCRIPTION';
// export const OPERATOR_EVENT_SUBSCRIPTION = 'OPERATOR_EVENT_SUBSCRIPTION';
export const MESHPLAY_CONTROLLER_SUBSCRIPTION = 'MESHPLAY_CONTROLLER_SUBSCRIPTION';

export const fnMapping = {
  // MESHSYNC_EVENT_SUBSCRIPTION : {
  //   eventName : "listenToMeshSyncEvents",
  //   comparatorFn : isMeshSyncSubscriptionDataUpdated,
  //   subscriptionFn : subscribeMeshSyncStatusEvents,
  //   mergeFn : mergeMeshSyncSubscription
  // },
  MESHPLAY_CONTROLLER_SUBSCRIPTION: {
    eventName: 'subscribeMeshplayControllersStatus',
    subscriptionFn: subscribeMeshplayControllersStatus,
    mergeFn: mergeMeshplayController,
    comparatorFn: isMeshplayControllerStateSubscriptionDataUpdated,
  },
  // OPERATOR_EVENT_SUBSCRIPTION: {
  //   eventName: 'operator',
  //   comparatorFn: isOperatorStateSubscriptionDataUpdated,
  //   subscriptionFn: subscribeOperatorStatusEvents,
  //   mergeFn: mergeOperatorStateSubscription,
  // },
};

export function isControllerObjectEqual(oldController, newController) {
  return (
    newController.contextId === oldController.contextId &&
    newController.controller === oldController.controller
  );
}
