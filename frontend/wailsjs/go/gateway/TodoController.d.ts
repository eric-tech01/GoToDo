// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT
import {todoModel} from '../models';
import {context} from '../models';

export function Delete(arg1:number):Promise<string>;

export function Finish(arg1:number):Promise<string>;

export function Load():Promise<todoModel.TodoList>;

export function Save(arg1:todoModel.Todo):Promise<todoModel.Todo>;

export function Startup(arg1:context.Context):Promise<void>;
