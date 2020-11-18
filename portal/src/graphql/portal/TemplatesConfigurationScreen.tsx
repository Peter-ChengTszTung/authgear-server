import React, {
  createContext,
  useCallback,
  useContext,
  useEffect,
  useMemo,
  useState,
} from "react";
import cn from "classnames";
import { useNavigate, useParams, useSearchParams } from "react-router-dom";
import produce from "immer";
import { Pivot, PivotItem, Text } from "@fluentui/react";
import { Context, FormattedMessage } from "@oursky/react-messageformat";

import { ModifiedIndicatorWrapper } from "../../ModifiedIndicatorPortal";
import ShowLoading from "../../ShowLoading";
import ShowError from "../../ShowError";
import TemplateLocaleManagement from "./TemplateLocaleManagement";
import ForgotPasswordTemplatesSettings from "./ForgotPasswordTemplatesSettings";
import PasswordlessAuthenticatorTemplatesSettings from "./PasswordlessAuthenticatorTemplatesSettings";
import { useAppConfigQuery } from "./query/appConfigQuery";
import { useAppTemplatesQuery } from "./query/appTemplatesQuery";
import { useUpdateAppTemplatesMutation } from "./mutations/updateAppTemplatesMutation";
import { useUpdateAppConfigMutation } from "./mutations/updateAppConfigMutation";
import { PortalAPIAppConfig } from "../../types";
import { usePivotNavigation } from "../../hook/usePivot";
import {
  DEFAULT_TEMPLATE_LOCALE,
  TemplateLocale,
  templatePaths,
} from "../../templates";
import { useGenericError } from "../../error/useGenericError";

import styles from "./TemplatesConfigurationScreen.module.scss";

interface AppConfigContextValue {
  effectiveAppConfig: PortalAPIAppConfig | null;
  rawAppConfig: PortalAPIAppConfig | null;
}

const FORGOT_PASSWORD_PIVOT_KEY = "forgot_password";
const PASSWORDLESS_AUTHENTICATOR_PIVOT_KEY = "passwordless_authenticator";

const AppConfigContext = createContext<AppConfigContextValue>({
  effectiveAppConfig: null,
  rawAppConfig: null,
});

const TemplatesConfiguration: React.FC = function TemplatesConfiguration() {
  const { renderToString } = useContext(Context);
  const { appID } = useParams();
  const navigate = useNavigate();
  const [searchParams] = useSearchParams();
  const { effectiveAppConfig, rawAppConfig } = useContext(AppConfigContext);

  const initialDefaultTemplateLocale = useMemo(() => {
    return (
      effectiveAppConfig?.localization?.fallback_language ??
      DEFAULT_TEMPLATE_LOCALE
    );
  }, [effectiveAppConfig]);

  const [remountIdentifier, setRemountIdentifier] = useState(0);
  const [defaultTemplateLocale, setDefaultTemplateLocale] = useState<
    TemplateLocale
  >(initialDefaultTemplateLocale);

  const [pendingTemplateLocales, setPendingTemplateLocales] = useState<
    TemplateLocale[]
  >([]);

  const templateLocale = useMemo(() => {
    const paramLocale = searchParams.get("locale");
    if (paramLocale != null && paramLocale.trim() !== "") {
      return paramLocale;
    }
    return defaultTemplateLocale;
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [searchParams]);

  const {
    templates,
    templateLocales: configuredLocales,
    loading: loadingTemplates,
    error: loadTemplatesError,
    refetch: refetchTemplates,
  } = useAppTemplatesQuery(appID, templateLocale, ...templatePaths);

  const {
    updateAppTemplates,
    loading: updatingTemplates,
    error: updateTemplatesError,
    resetError: resetUpdateTemplatesError,
  } = useUpdateAppTemplatesMutation(appID);

  const {
    updateAppConfig,
    loading: updatingAppConfig,
    error: updateAppConfigError,
    resetError: resetUpdateAppConfigError,
  } = useUpdateAppConfigMutation(appID);

  const saveDefaultTemplateLocale = useCallback(
    (defaultTemplateLocale: TemplateLocale) => {
      if (rawAppConfig == null) {
        return;
      }
      const newAppConfig = produce(rawAppConfig, (draftConfig) => {
        draftConfig.localization = draftConfig.localization ?? {};
        draftConfig.localization.fallback_language = defaultTemplateLocale;
      });

      updateAppConfig(newAppConfig).catch(() => {});
    },
    [rawAppConfig, updateAppConfig]
  );

  const refresh = useCallback(() => {
    setRemountIdentifier((prev) => prev + 1);
  }, []);

  const resetError = useCallback(() => {
    resetUpdateTemplatesError();
    resetUpdateAppConfigError();
  }, [resetUpdateTemplatesError, resetUpdateAppConfigError]);

  const resetForm = useCallback(() => {
    refresh();
    resetError();
  }, [resetError, refresh]);

  const onTemplateLocaleSelected = useCallback(
    (locale: TemplateLocale) => {
      navigate(`?locale=${locale}${window.location.hash}`);
    },
    [navigate]
  );

  // NOTE: handle invalid locale key in query string
  useEffect(() => {
    const localeList = configuredLocales.concat(pendingTemplateLocales);
    if (
      !localeList.includes(templateLocale) &&
      localeList.includes(defaultTemplateLocale)
    ) {
      onTemplateLocaleSelected(defaultTemplateLocale);
    }
  }, [
    onTemplateLocaleSelected,
    configuredLocales,
    pendingTemplateLocales,
    defaultTemplateLocale,
    templateLocale,
  ]);

  useEffect(() => {
    refresh();
  }, [templateLocale, refresh]);

  useEffect(() => {
    for (const configuredLocale of configuredLocales) {
      if (pendingTemplateLocales.includes(configuredLocale)) {
        setPendingTemplateLocales((prev) =>
          prev.filter((locale) => locale !== configuredLocale)
        );
      }
    }
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [configuredLocales]);

  useEffect(() => {
    if (!loadingTemplates) {
      refresh();
    }
  }, [loadingTemplates, refresh]);

  const { selectedKey, onLinkClick } = usePivotNavigation(
    [FORGOT_PASSWORD_PIVOT_KEY, PASSWORDLESS_AUTHENTICATOR_PIVOT_KEY],
    resetError
  );

  const { unrecognizedError: unrecognizedLoadTemplateError } = useGenericError(
    loadTemplatesError,
    [],
    [
      {
        reason: "ResourceNotFound",
        // NOTE: error message unused
        errorMessageID: "generic-error.unknown-error",
      },
    ]
  );

  return (
    <main
      className={cn(styles.root, {
        [styles.loading]: updatingTemplates,
      })}
    >
      {updateTemplatesError && <ShowError error={updateTemplatesError} />}
      {updateAppConfigError && <ShowError error={updateAppConfigError} />}
      {unrecognizedLoadTemplateError && (
        <ShowError error={loadTemplatesError} onRetry={refetchTemplates} />
      )}
      <ModifiedIndicatorWrapper className={styles.screen}>
        <Text className={styles.screenHeaderText} as="h1">
          <FormattedMessage id="TemplatesConfigurationScreen.title" />
        </Text>
        <TemplateLocaleManagement
          key={remountIdentifier}
          configuredTemplateLocales={configuredLocales}
          templateLocale={templateLocale}
          initialDefaultTemplateLocale={initialDefaultTemplateLocale}
          defaultTemplateLocale={defaultTemplateLocale}
          onTemplateLocaleSelected={onTemplateLocaleSelected}
          onDefaultTemplateLocaleSelected={setDefaultTemplateLocale}
          pendingTemplateLocales={pendingTemplateLocales}
          onPendingTemplateLocalesChange={setPendingTemplateLocales}
          saveDefaultTemplateLocale={saveDefaultTemplateLocale}
          updatingAppConfig={updatingAppConfig}
        />
        {loadingTemplates && <ShowLoading />}
        <Pivot
          hidden={loadingTemplates || unrecognizedLoadTemplateError != null}
          onLinkClick={onLinkClick}
          selectedKey={selectedKey}
        >
          <PivotItem
            headerText={renderToString(
              "TemplatesConfigurationScreen.forgot-password.title"
            )}
            itemKey={FORGOT_PASSWORD_PIVOT_KEY}
          >
            <ForgotPasswordTemplatesSettings
              key={remountIdentifier}
              templates={templates}
              templateLocale={templateLocale}
              updateTemplates={updateAppTemplates}
              updatingTemplates={updatingTemplates}
              resetForm={resetForm}
            />
          </PivotItem>
          <PivotItem
            headerText={renderToString(
              "TemplatesConfigurationScreen.passwordless-authenticator.title"
            )}
            itemKey={PASSWORDLESS_AUTHENTICATOR_PIVOT_KEY}
          >
            <PasswordlessAuthenticatorTemplatesSettings
              key={remountIdentifier}
              templates={templates}
              templateLocale={templateLocale}
              updateTemplates={updateAppTemplates}
              updatingTemplates={updatingTemplates}
              resetForm={resetForm}
            />
          </PivotItem>
        </Pivot>
      </ModifiedIndicatorWrapper>
    </main>
  );
};

const TemplatesConfigurationScreen: React.FC = function TemplatesConfigurationScreen() {
  const { appID } = useParams();
  const {
    effectiveAppConfig,
    rawAppConfig,
    loading,
    error,
    refetch,
  } = useAppConfigQuery(appID);

  const appConfigContextValue = useMemo(() => {
    return {
      effectiveAppConfig,
      rawAppConfig,
    };
  }, [effectiveAppConfig, rawAppConfig]);

  if (loading) {
    return <ShowLoading />;
  }

  if (error) {
    return <ShowError error={error} onRetry={refetch} />;
  }

  return (
    <AppConfigContext.Provider value={appConfigContextValue}>
      <TemplatesConfiguration />
    </AppConfigContext.Provider>
  );
};

export default TemplatesConfigurationScreen;