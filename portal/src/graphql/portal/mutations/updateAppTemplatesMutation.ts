import React from "react";
import { useMutation, gql } from "@apollo/client";

import { client } from "../apollo";
import { AppResourceUpdate } from "../__generated__/globalTypes";
import {
  UpdateAppTemplatesMutation,
  UpdateAppTemplatesMutationVariables,
} from "./__generated__/UpdateAppTemplatesMutation";
import { PortalAPIApp } from "../../../types";

const updateAppTemplatesMutation = gql`
  mutation UpdateAppTemplatesMutation(
    $appID: ID!
    $updates: [AppResourceUpdate!]!
    $paths: [String!]!
  ) {
    updateAppResources(input: { appID: $appID, updates: $updates }) {
      app {
        id
        resources(paths: $paths) {
          path
          effectiveData
        }
      }
    }
  }
`;

export type AppTemplatesUpdater<TemplatePath extends string> = (
  updateTemplates: {
    [path in TemplatePath]?: string | null;
  }
) => Promise<PortalAPIApp | null>;

export function useUpdateAppTemplatesMutation<TemplatePath extends string>(
  appID: string
): {
  updateAppTemplates: AppTemplatesUpdater<TemplatePath>;
  loading: boolean;
  error: unknown;
} {
  const [mutationFunction, { error, loading }] = useMutation<
    UpdateAppTemplatesMutation,
    UpdateAppTemplatesMutationVariables
  >(updateAppTemplatesMutation, { client });
  const updateAppTemplates = React.useCallback(
    async (updateTemplates: { [path in TemplatePath]?: string | null }) => {
      const paths: string[] = [];
      const updates: AppResourceUpdate[] = [];
      for (const [path, data] of Object.entries(updateTemplates)) {
        if (data === undefined) {
          continue;
        }
        paths.push(path);
        updates.push({ path, data: data as string | null });
      }

      const result = await mutationFunction({
        variables: {
          appID,
          paths,
          updates,
        },
      });
      return result.data?.updateAppResources.app ?? null;
    },
    [appID, mutationFunction]
  );
  return { updateAppTemplates, error, loading };
}
