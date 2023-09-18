import { useEffect, useState } from 'react';
import { Subscription } from 'rxjs';
import { ConAuthorization } from '../../models/types';
import { conAuthorizationRef$ } from '../observable';

export const useAuthorizationHook = (id?: string) => {
    const [conAuthorization, setConAuthorization] = useState<ConAuthorization>();
    const [loading, setLoading] = useState<boolean>(true);
    useEffect(() => {
        let conAuthorizationObservable: Subscription;
        if (id) {
            conAuthorizationObservable = conAuthorizationRef$(id).subscribe((userSetting) => {
                setConAuthorization(userSetting as ConAuthorization);
                setLoading(false);
            });
        }
        return () => {
            if (conAuthorization) {
                conAuthorizationObservable.unsubscribe();
            }
        };
    }, [id]);

    return { conAuthorization, loading };
};
