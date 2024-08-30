import { workerfyActor } from '@khulnasoft/sistent';
import { schemaValidatorMachine } from './schemaValidator';

console.log('Workerfying schemaValidatorMachine');

workerfyActor(schemaValidatorMachine);
